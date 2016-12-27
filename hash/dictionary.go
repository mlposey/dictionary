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
// TODO: fnv hashes
type Dictionary struct {
	tableA []*Pair
	tableB []*Pair

	// The number of possible hash values
	// tableSize should always be a power of 2. This allows us to avoid
	// mods in favor of ands.
	tableSize int

	// The number of items in the Dictionary
	Size int

	hashConstant int
}

// NewDictionary creates an empty Dictionary with an optional table size.
//
// tableSize should be a power of two. If it is not, a higher tableSize
// will be chosen which is a power of two.
func NewDictionary(tableSize ...int) *Dictionary {
	finalSize := 8
	if len(tableSize) != 0 {
		// Make capacity a power of two.
		finalSize = int(math.Exp2(math.Ceil(math.Log2(float64(tableSize[0])))))
	}
	return &Dictionary{
		tableA:       make([]*Pair, finalSize),
		tableB:       make([]*Pair, finalSize),
		tableSize:    finalSize,
		hashConstant: rand.Intn(100) + 1,
	}
}

// Insert adds an object to the Dictionary.
//
// If key is not an int or string, it should implement the Hashable interface.
// Its hash function should not perform a modulus-type bounding operation since
// this is done internally. All other functions which have key parameters also
// make this assumption.
func (d *Dictionary) Insert(key interface{}, object interface{}) error {
	if d.Size == d.tableSize {
		d.resize()
	}

	// This is the maximum number of loop iterations that will occur
	// before a rehash is triggered. The value is taken from "Efficient
	// Hash Probes on Modern Processors" by Kenneth A. Ross.
	const kMaxLoop = 1000

	for i := 0; i < kMaxLoop; i++ {
		index := d.hashKeyA(key) & (d.tableSize - 1)
		if d.tableA[index] == nil {
			d.tableA[index] = &Pair{Key: key, Value: object}
			d.Size++
			return nil
		}

		index = d.hashKeyB(key) & (d.tableSize - 1)
		if d.tableB[index] == nil {
			d.tableB[index] = &Pair{Key: key, Value: object}
			d.Size++
			return nil
		}

		// Replace the old tableB object and use the next iteration
		// to find its home.
		old := d.tableB[index]
		d.tableB[index] = &Pair{Key: key, Value: object}
		key = old.Key
		object = old.Value
	}
	d.rehash()
	d.Insert(key, object)
	return nil
}

// Get retrieves an object from the Dictionary.
//
// If the provided key is not in the Dictionary, nil is returned.
//
// Get will always run in O(1) time.
func (d *Dictionary) Get(key interface{}) interface{} {
	// Check tableA
	index := d.hashKeyA(key) & (d.tableSize - 1)
	if d.tableA[index] != nil && d.tableA[index].Key == key {
		return d.tableA[index].Value
	}

	// Check tableB
	index = d.hashKeyB(key) & (d.tableSize - 1)
	if d.tableB[index] != nil && d.tableB[index].Key == key {
		return d.tableB[index].Value
	}

	return nil
}

// Remove removes an object from the Dictionary.
//
// Returns:
//	(1) nil if the operation was successful
//	(2) an error if the key is not in the Dictionary
func (d *Dictionary) Remove(key interface{}) error {
	// Check tableA
	index := d.hashKeyA(key) & (d.tableSize - 1)
	if d.tableA[index] != nil && d.tableA[index].Key == key {
		d.tableA[index] = nil
		d.Size--
		return nil
	}

	// Check tableB
	index = d.hashKeyB(key) & (d.tableSize - 1)
	if d.tableB[index] != nil && d.tableB[index].Key == key {
		d.tableB[index].Value = nil
		d.Size--
		return nil
	}
	return errors.New("Key not found")
}

// hashKey hashes the key to a uint.
//
// This value is not guaranteed to be in the range [0, tableSize).
//
// Universal hash procedures are provided for both string and uint. If
// key is not either of those, it is assumed to implement the Hashable
// interface which contains a Hash() function.
func (d *Dictionary) hashKeyA(key interface{}) int {
	switch key.(type) {
	case string:
		// TODO: universal string hash
		var res int
		kLen := len(key.(string))
		for i, char := range key.(string) {
			res += int(char) * int(math.Pow(float64(d.hashConstant), float64(kLen-i)))
		}
		return res
	case int:
		// TODO: universal int hash
		return d.hashConstant * key.(int)
	default:
		return d.hashConstant * key.(Hashable).Hash()
	}
}

// hashKey hashes a result of hashKeyA.
//
// This value is not guaranteed to be in the range [0, tableSize)
func (d *Dictionary) hashKeyB(key interface{}) int {
	return d.hashKeyA(key) / d.tableSize
}

// resize increases the size of the table and rehashes all keys.
//
// The new table size will be 2^(ceil(log2(oldSize * 2))).
func (d *Dictionary) resize(size ...int) {
	var newSize int
	if len(size) != 0 {
		newSize = size[0]
	} else {
		newSize = int(math.Exp2(math.Ceil(math.Log2(float64(d.tableSize * 2)))))

	}
	a, b := make([]*Pair, newSize), make([]*Pair, newSize)

	// Give d the new tables.
	d.tableSize = newSize
	d.tableA, a = a, d.tableA
	d.tableB, b = b, d.tableB

	// Insert the old values from a and b into the new tables.
	for i := range a {
		if a[i] != nil {
			d.Insert(a[i].Key, a[i].Value)
		}
		if b[i] != nil {
			d.Insert(b[i].Key, b[i].Value)
		}
	}
}

// rehash creates new hash functions and rebuilds the Dictionary.
func (d *Dictionary) rehash() {
	d.hashConstant = rand.Intn(100) + 1

	// Use the new hash function to determine key locations.
	d.resize(d.tableSize)
}

// Hashable is an interface for objects which can compute their hash values.
type Hashable interface {
	Hash() int
}

// Pair is an ordered pair of a key and an object identified by that key.
type Pair struct {
	Key   interface{}
	Value interface{}
}
