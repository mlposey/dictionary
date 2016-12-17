package heap

import "testing"

func TestNewHeap(t *testing.T) {
	h := NewHeap(func(a, b interface{}) int8 {
		if a.(int) < b.(int) {
			return -1
		} else if a.(int) == b.(int) {
			return 0
		} else {
			return 1
		}
	})

	if h.Capacity != 20 {
		t.Error("Wrong initial capacity of Heap")
	}
	if len(h.items) != int(h.Capacity) {
		t.Error("Could not initialize heap slice")
	}
}
