package heap

import "testing"

func heapComparator(a, b interface{}) int8 {
	if a.(int) < b.(int) {
		return -1
	} else if a.(int) == b.(int) {
		return 0
	} else {
		return 1
	}
}

// Test NewHeap
func TestNewHeap(t *testing.T) {
	h := NewHeap(heapComparator, false)

	if h.Capacity != kDefaultHeapCapacity {
		t.Error("Wrong initial capacity of Heap")
	}
	if len(h.items) != int(h.Capacity) {
		t.Error("Could not initialize heap slice")
	}
}

// Test *Heap.Insert
func TestHeap_Insert(t *testing.T) {
	h := NewHeap(heapComparator, false, 2)

	h.Insert(3)
	h.Insert(5)
	h.Insert(9)

	if h.items[0] != 9 {
		t.Error("Expected", 9, "got", h.items[0])
	}
}

// Test *Heap.Insert
//
// Heap should allow objects of duplicate weights.
func TestHeap_Insert_Duplicate(t *testing.T) {
	h := NewMinHeap(heapComparator)

	h.Insert(3)
	h.Insert(3)

	if h.Size != 2 {
		t.Error("Could not insert objects with duplicate weights")
	}
}

func TestHeap_Remove(t *testing.T) {
	h := NewHeap(heapComparator, false, 2)

	h.Insert(3)
	h.Insert(5)
	h.Insert(1)

	root := h.Remove()
	if root != 5 {
		t.Error("Expected", 5, "got", root)
	}
}

// Test *Heap.Remove
//
// Calling remove on an empty heap should return nil.
func TestHeap_Remove_Empty(t *testing.T) {
	h := NewMinHeap(heapComparator)

	if h.Remove() != nil {
		t.Error("Encountered problem removing from empty Heap")
	}
}
