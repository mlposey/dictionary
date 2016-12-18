package heap

// Heap implements a binary heap.
//
// Greater -
//	This required function compares two heap objects. The return value
//	should be:
//		-1 if arg1 < arg2
//		0  if arg1 == arg2
//		1  if arg1 > arg2
type Heap struct {
	// Compare two heap objects
	Greater func(interface{}, interface{}) int8

	// The max number of items Heap can hold before resizing
	Capacity uint

	// The number of items in the Heap
	Size uint

	items []interface{}
}

// NewHeap creates a new *Heap.
//
// minHeap - true if Heap should be a min-heap; false otherwise
//
// capacity - the initial capacity of the Heap
//	  If not supplied, a default capacity is chosen.
//
// see Heap doc for valid greater definitions.
func NewHeap(greater func(interface{}, interface{}) int8, minHeap bool, capacity ...uint) *Heap {
	var cap uint
	if len(capacity) == 0 {
		cap = 20
	} else {
		cap = capacity[0]
	}
	comp := greater
	if minHeap {
		comp = func(a, b interface{}) int8 {
			return -greater(a, b)
		}
	}
	return &Heap{
		Greater:  comp,
		Capacity: cap,
		items:    make([]interface{}, cap),
	}
}

// NewMinHeap creates and returns a new *Heap with the min-heap property.
//
// The root of a min-heap always contains the smallest key. Thus, Remove()
// will return the smallest, in an abstract sense, object.
//
// capacity - the initial capacity of the Heap
//	  If not supplied, a default capacity is chosen.
//
// see Heap doc for valid greater definitions.
func NewMinHeap(greater func(interface{}, interface{}) int8, capacity ...uint) *Heap {
	return NewHeap(greater, true, capacity)
}

// NewMaxHeap creates and returns a new *Heap with the max-heap property.
//
// The root of a max-heap always contains the largest key. Thus, Remove()
// will return the largest, in an abstract sense, object.
//
// capacity - the initial capacity of the Heap
//	  If not supplied, a default capacity is chosen.
//
// see Heap doc for valid greater definitions.
func NewMaxHeap(greater func(interface{}, interface{}) int8, capacity ...uint) *Heap {
	return NewHeap(greater, false, capacity)
}

// Insert adds an object to the Heap.
func (h *Heap) Insert(object interface{}) {
	h.items[h.Size] = object
	parent := (h.Size - 1) / 2
	for i := h.Size; i != 0 && h.Greater(h.items[i], h.items[parent]) == 1; i = parent {
		parent = (i - 1) / 2
		h.items[parent], h.items[i] = h.items[i], h.items[parent]
	}
	h.Size++

	// Expand heap capacity if necessary
	if h.Size == h.Capacity {
		h.Capacity = h.Capacity * 2
		newSlice := make([]interface{}, h.Capacity)
		copy(newSlice, h.items)
		h.items = newSlice
	}
}

// Remove removes and returns the root from the heap.
//
// Returns nil if the Heap is empty.
func (h *Heap) Remove() interface{} {
	if h.Size == 0 {
		return nil
	}
	root := h.items[0]
	h.Size--
	h.items[0] = h.items[h.Size]

	// Reheap the heap
	var big uint = 1
	var i uint = 0
	for big < h.Size {
		if big < h.Size-1 &&
			h.Greater(h.items[big+1], h.items[big]) == 1 {
			big++
		}
		if h.Greater(h.items[big], h.items[i]) == 1 {
			h.items[i], h.items[big] = h.items[big], h.items[i]
			i = big
			big = 2*i + 1
		} else {
			break
		}
	}

	return root
}
