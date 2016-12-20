package hash

type Hashable interface {
	Hash() uint
}

type Dictionary struct {
	table []uint
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
