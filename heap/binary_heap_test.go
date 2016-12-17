package heap

import "testing"

// Test NewHeap
func TestNewHeap(t *testing.T) {
	h := NewHeap(func(a, b interface{}) int8 {
		if a.(int) < b.(int) {
			return -1
		} else if a.(int) == b.(int) {
			return 0
		} else {
			return 1
		}
	}, false)

	if h.Capacity != 20 {
		t.Error("Wrong initial capacity of Heap")
	}
	if len(h.items) != int(h.Capacity) {
		t.Error("Could not initialize heap slice")
	}
}

// Test *Heap.Insert
func TestHeap_Insert(t *testing.T) {
	h := NewHeap(func(a, b interface{}) int8 {
		if a.(int) < b.(int) {
			return -1
		} else if a.(int) == b.(int) {
			return 0
		} else {
			return 1
		}
	}, false, 2)

	h.Insert(3)
	h.Insert(5)
	h.Insert(9)

	if h.items[0] != 9 {
		t.Error("Expected", 9, "got", h.items[0])
	}
}
