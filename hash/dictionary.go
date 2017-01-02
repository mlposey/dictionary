package hash

import (
	"errors"
)

// Each slot in a table holds a Bucket. Each Bucket can hold no more than
// kBucketCapacity *Pairs.
type Bucket []*Pair

// Dictionary is a hash table for large volume data sets.
//
// It currently supports int, string, and Hashable key types. Types which
// implement Hashable should not bound the result of Hash; this is done
// internally. When using Dictionary, keep in mind that: (1) It currently does
// not support concurrent use and (2) For small data sets, the built-in map
// type may be a better option.
//
// Complexity:
//	Insert - Amortized O(1)
//	Get - Worst-case O(1)
//	Remove - Worst-case O(1)
//
// TODO: concurrency
type Dictionary struct {
	tables [kTableCount][]Bucket
	hashes [kTableCount]Hasher

	// The number of possible hash values
	tableSize int

	// The number of items in the Dictionary
	Size int
}

const (
	// The number of tables in the Dictionary
	// This is also the number of hash functions in use.
	kTableCount = 2

	// As per "Efficient Hash Probes on Modern Processors" by Kenneth A.
	// Ross, a bucket capacity of 4 with two tables will provide a load
	// factor of 0.976.
	kBucketCapacity = 4
)

// NewDictionary creates an empty Dictionary with an optional table size.
func NewDictionary(tableSize ...int) *Dictionary {
	// Acquire a capacity for the tables which is a power of two.
	capacity := 8
	if len(tableSize) != 0 {
		capacity = tableSize[0]
	}

	dict := &Dictionary{tableSize: capacity}
	for i := range dict.tables {
		dict.tables[i] = make([]Bucket, capacity)
	}
	return dict
}

// Insert adds an object to the Dictionary.
//
// If key is not an int or string, it should implement the Hashable interface.
// Its hash function should not perform a modulus-type bounding operation since
// this is done internally. All other functions which have key parameters also
// make this assumption.
func (d *Dictionary) Insert(key interface{}, object interface{}) {
	// If this is the first insertion, determine if an in-house
	// hash function can be used for this and future hashing.
	if d.hashes[0] == nil {
		switch key.(type) {
		case string:
			d.generateStringHashers()
		case int32:
			d.generateInt32Hashers()
		}
	}

	if d.Size == d.tableSize {
		d.resize()
	}

	// This is the maximum number of loop iterations that will occur
	// before a rehash is triggered. The value is taken from "Efficient
	// Hash Probes on Modern Processors" by Kenneth A. Ross.
	const kMaxLoop = 1000

	for {
		for i := 0; i < kMaxLoop; i++ {
			var index, t int
			// Try each table the key could be located in.
			for ; t < kTableCount; t++ {
				index = d.hashes[t].Hash(key) % d.tableSize

				// Is there even room in this bucket?
				if len(d.tables[t][index]) != kBucketCapacity {
					d.tables[t][index] = append(d.tables[t][index], &Pair{
						Key:   key,
						Value: object,
					})
					d.Size++
					return
				}
			}
			t--

			// Replace the object at the last checked index and attempt
			// to re-home it in the next iteration.
			old := d.tables[t][index][0]
			d.tables[t][index][0] = &Pair{Key: key, Value: object}
			key = old.Key
			object = old.Value
		}
		d.rehash()
		// TODO: Insert may loop forever.
	}

}

// Get retrieves an object from the Dictionary.
//
// If the provided key is not in the Dictionary, nil is returned.
//
// Get will always run in O(1) time.
func (d *Dictionary) Get(key interface{}) interface{} {
	var index int

	// Check each possible location for key.
	for i := range d.tables {
		index = d.hashes[i].Hash(key) % d.tableSize
		for b := range d.tables[i][index] {
			if d.tables[i][index][b].Key == key {
				return d.tables[i][index][b].Value
			}
		}
	}
	return nil
}

// Remove removes an object from the Dictionary.
//
// Returns:
//	(1) nil if the operation was successful
//	(2) an error if the key is not in the Dictionary
func (d *Dictionary) Remove(key interface{}) error {
	var index int

	// Check each possible location for key.
	for i := range d.tables {
		index = d.hashes[i].Hash(key) % d.tableSize
		for b := range d.tables[i][index] {
			if d.tables[i][index][b].Key == key {
				d.tables[i][index] = append(d.tables[i][index][:b], d.tables[i][index][b+1:]...)
				d.Size--
				return nil
			}
		}
	}
	return errors.New("Key not found")
}

// resize increases the size of the table and rehashes all keys.
func (d *Dictionary) resize(size ...int) {
	if len(size) != 0 {
		d.tableSize = size[0]
	} else {
		d.tableSize *= 2
	}

	// Create new tables with the increased size and swap them into d.
	var oldTables [kTableCount][]Bucket
	for i := range oldTables {
		oldTables[i] = make([]Bucket, d.tableSize)
		oldTables[i], d.tables[i] = d.tables[i], oldTables[i]
	}
	d.Size = 0

	// Insert the keys from the old tables into the new ones.
	for t := range oldTables {
		for b := range oldTables[t] {
			for i := range oldTables[t][b] {
				if oldTables[t][b][i] != nil {
					d.Insert(oldTables[t][b][i].Key,
						oldTables[t][b][i].Value)
				}
			}
		}
	}
}

func (d *Dictionary) generateStringHashers() {
	for i := range d.hashes {
		d.hashes[i] = NewStringHasher()
	}
}

func (d *Dictionary) generateInt32Hashers() {
	for i := range d.hashes {
		d.hashes[i] = NewIntHasher()
	}
}

// rehash creates new hash functions and rebuilds the Dictionary.
func (d *Dictionary) rehash() {
	for i := range d.hashes {
		d.hashes[i].Reseed()
	}

	// Use the new hash function to determine key locations.
	d.resize(d.tableSize)
}

// Pair is an ordered pair of a key and an object identified by that key.
type Pair struct {
	Key   interface{}
	Value interface{}
}
