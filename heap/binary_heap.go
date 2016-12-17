package heap

// Heap implements a binary heap.
//
// KeyCompare -
//	This required function compares two heap objects. The return value
//	should be:
//		-1 if arg1 < arg2
//		0  if arg1 == arg2
//		1  if arg1 > arg2
type Heap struct {
	// Compare two heap objects
	KeyCompare func(interface{}, interface{}) int8

	// The max number of items Heap can hold before resizing
	Capacity uint

	// The number of items in the Heap
	Size uint

	items []interface{}
}
