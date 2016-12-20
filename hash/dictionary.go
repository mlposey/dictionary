package hash

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
}

// hashKey hashes the key to a uint.
//
// Universal hash procedures are provided for both string and uint. If
// key is not either of those, it is assumed to implement the Hashable
// interface which contains a Hash() function.
func hashKey(key interface{}) uint {
	switch key.(type) {
	case string:
		// TODO: universal string hash
		return 0
	case uint:
		// TODO: universal int hash
		return 0
	default:
		return key.(Hashable).Hash()
	}
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

// Hashable is an interface for objects which can compute their hash values.
type Hashable interface {
	Hash() uint
}
