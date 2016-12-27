package hash

import (
	"errors"
	"math"
	"math/rand"
)

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
// TODO: buckets
type Dictionary struct {
	tables [kTableCount][]*Pair
	hashes [kTableCount]HashFunction

	// The number of possible hash values
	tableSize int

	// The number of items in the Dictionary
	Size int
}

const (
	// The number of tables in the Dictionary
	// This is also the number of hash functions in use.
	kTableCount = 2
)

// NewDictionary creates an empty Dictionary with an optional table size.
func NewDictionary(tableSize ...int) *Dictionary {
	// Acquire a capacity for the tables which is a power of two.
	capacity := 8
	if len(tableSize) != 0 {
		capacity = tableSize[0]
	}

	// Initialize the Dictionary.
	dict := &Dictionary{tableSize: capacity}
	dict.generateHashFunctions()
	for i := range dict.tables {
		dict.tables[i] = make([]*Pair, capacity)
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
	if d.Size == d.tableSize {
		d.resize()
	}

	// This is the maximum number of loop iterations that will occur
	// before a rehash is triggered. The value is taken from "Efficient
	// Hash Probes on Modern Processors" by Kenneth A. Ross.
	const kMaxLoop = 1000

	for i := 0; i < kMaxLoop; i++ {
		var index, t int
		// Try each table the key could be located in.
		for ; t < kTableCount; t++ {
			index = d.hashes[t](key) % d.tableSize
			if d.tables[t][index] == nil {
				d.tables[t][index] = &Pair{
					Key:   key,
					Value: object,
				}
				d.Size++
				return
			}
		}
		t--

		// Replace the object at the last checked index and attempt
		// to re-home it in the next iteration.
		old := d.tables[t][index]
		d.tables[t][index] = &Pair{Key: key, Value: object}
		key = old.Key
		object = old.Value
	}
	d.rehash()
	d.Insert(key, object)
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
		index = d.hashes[i](key) % d.tableSize
		if d.tables[i][index] != nil && d.tables[i][index].Key == key {
			return d.tables[i][index].Value
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
		index = d.hashes[i](key) % d.tableSize
		if d.tables[i][index] != nil && d.tables[i][index].Key == key {
			d.tables[i][index] = nil
			d.Size--
			return nil
		}
	}
	return errors.New("Key not found")
}

// resize increases the size of the table and rehashes all keys.
//
// The new table size will be 2^(ceil(log2(oldSize * 2))).
func (d *Dictionary) resize(size ...int) {
	if len(size) != 0 {
		d.tableSize = size[0]
	} else {
		d.tableSize *= 2
	}

	// Create new tables with the increased size and swap them into d.
	var oldTables [kTableCount][]*Pair
	for i := range oldTables {
		oldTables[i] = make([]*Pair, d.tableSize)
		oldTables[i], d.tables[i] = d.tables[i], oldTables[i]
	}
	d.Size = 0

	// Insert the keys from the old tables into the new ones.
	for t := range oldTables {
		for i := range oldTables[t] {
			if oldTables[t][i] != nil {
				d.Insert(oldTables[t][i].Key, oldTables[t][i].Value)
			}
		}
	}
}

// generateHashFunctions creates kTableCount hash functions for the Dictionary.
func (d *Dictionary) generateHashFunctions() {
	// Make a hash function that uses a certain unique factor.
	generate := func(factor int) HashFunction {
		return func(key interface{}) int {
			switch key.(type) {
			case string:
				// TODO: universal string hash
				var res int
				kLen := len(key.(string))
				for i, char := range key.(string) {
					res += int(char) * int(math.Pow(float64(factor), float64(kLen-i)))
				}
				return int(math.Abs(float64(res))) // for overflows
			case int:
				// TODO: universal int hash
				return factor * key.(int)
			default:
				return factor * key.(Hashable).Hash()
			}
		}
	}

	// Make hash functions for each table.
	for i := range d.hashes {
		if i == 0 || i%2 == 0 {
			d.hashes[i] = generate(rand.Intn(100) + 1)
		} else {
			d.hashes[i] = func(key interface{}) int {
				return d.hashes[i-1](key) / d.tableSize
			}
		}
	}
}

// rehash creates new hash functions and rebuilds the Dictionary.
func (d *Dictionary) rehash() {
	d.generateHashFunctions()

	// Use the new hash function to determine key locations.
	d.resize(d.tableSize)
}

// HashFunction is a function which turns a key into an integer.
type HashFunction func(key interface{}) int

// Hashable is an interface for objects which can compute their hash values.
type Hashable interface {
	Hash() int
}

// Pair is an ordered pair of a key and an object identified by that key.
type Pair struct {
	Key   interface{}
	Value interface{}
}
