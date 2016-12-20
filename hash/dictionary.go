package hash

type Hashable interface {
	Hash() uint
}

type Dictionary struct {
	table []uint
}

func (d *Dictionary) Insert(key Hashable, object interface{}) error {
	return nil
}

func (d *Dictionary) Get(key Hashable) interface{} {
	return nil
}

func (d *Dictionary) Remove(key Hashable) error {
	return nil
}
