package hash

import "math"

// Dictionary is a hash table for large volume data sets.
//
// If you expect the number of stored elements will be small, consider using
// a map, instead.
//
// TODO: Mention concurrency
// TODO: Mention Big-O estimates
// TODO: Mention Hashable interface
type Dictionary struct {
	table []uint

	// The number of possble hash values
	// tableSize should always be a power of 2. This allows us to avoid
	// mods in favor of ands.
	tableSize uint
}

// NewDictionary creates an empty Dictionary with an optional table size.
//
// tableSize should be a power of two. If it is not, a higher tableSize
// will be chosen which is a power of two.
func NewDictionary(tableSize ...uint) *Dictionary {
	var finalSize uint = 8
	if len(tableSize) != 0 {
		// Make capacity a power of two.
		finalSize = uint(math.Exp2(math.Ceil(math.Log2(float64(tableSize[0])))))
	}
	return &Dictionary{tableSize: finalSize}
}

func (d *Dictionary) Insert(key interface{}, object interface{}) error {
	return nil
}

func (d *Dictionary) Get(key interface{}) interface{} {
	return nil
}

func (d *Dictionary) Remove(key interface{}) error {
	return nil
}

// hashKey hashes the key to a uint.
//
// This value is not guaranteed to be in the range [0, tableSize).
//
// Universal hash procedures are provided for both string and uint. If
// key is not either of those, it is assumed to implement the Hashable
// interface which contains a Hash() function.
func (d *Dictionary) hashKeyA(key interface{}) uint {
	switch key.(type) {
	case string:
		// TODO: universal string hash
		var res uint = 13
		for _, char := range key.(string) {
			res *= 31 + uint(char)
		}
		return res
	case uint:
		// TODO: universal int hash
		return key.(uint)
	default:
		return key.(Hashable).Hash()
	}
}

// hashKey hashes a result of hashKeyA.
//
// This value is not guaranteed to be in the range [0, tableSize)
func (d *Dictionary) hashKeyB(hashA uint) uint {
	return hashA / d.tableSize
}

// Hashable is an interface for objects which can compute their hash values.
type Hashable interface {
	Hash() uint
}
