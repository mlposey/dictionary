package heap

// Heap implements a binary heap.
//
// Comparator -
//	This required function compares two heap objects. The return value
//	should be:
//		-1 if arg1 < arg2
//		0  if arg1 == arg2
//		1  if arg1 > arg2
type Heap struct {
	// Compare two heap objects
	Comparator func(interface{}, interface{}) int8

	// The max number of items Heap can hold before resizing
	Capacity uint

	// The number of items in the Heap
	Size uint

	items []interface{}
}

// NewHeap creates a new *Heap.
//
// size - the initial capacity of the Heap
//	  If not supplied, a default capacity is chosen.
//
// see Heap doc for valid comparator definitions.
func NewHeap(comparator func(interface{}, interface{}) int8, capacity ...uint) *Heap {
	var cap uint
	if len(capacity) == 0 {
		cap = 20
	} else {
		cap = capacity[0]
	}
	return &Heap{
		Comparator: comparator,
		Capacity:   cap,
		items:      make([]interface{}, cap),
	}
}
