package list

import "testing"

// Test *List.Insert
func TestList_Insert(t *testing.T) {
	l := &List{}

	l.Insert(3)
	if l.head == nil || l.tail == nil {
		t.Error("Nil head/tail after first List insert")
	}

	l.Insert(5)
	if l.head == l.tail {
		t.Error("Head does not change on List insert")
	}
	if l.head.next != l.tail {
		t.Error("Head pointer not changed on List insert")
	}
}

// Test *List.InsertEnd
func TestList_InsertEnd(t *testing.T) {
	l := &List{}

	l.InsertEnd(3)
	if l.head == nil || l.tail == nil {
		t.Error("Nil head/tail after first List InsertEnd")
	}

	l.InsertEnd(5)
	if l.head == l.tail {
		t.Error("Head does not change on List InsertEnd")
	}
	if l.tail.prev != l.head {
		t.Error("Tail pointer not changed on List InsertEnd")
	}
}

// Test *List.RemoveFront
func TestList_RemoveFront(t *testing.T) {
	l := &List{}

	if l.RemoveFront() != nil {
		t.Error("Non-nil return value on empty List")
	}

	testReturn := func(expect int) {
		if l.RemoveFront() != expect {
			t.Error("Wrong return value from *List.RemoveFront")
		}
	}

	l.Insert(3)
	testReturn(3)

	l.Insert(3)
	l.Insert(5)
	testReturn(5)
}

// Test *List.RemoveEnd
func TestList_RemoveEnd(t *testing.T) {
	l := &List{}

	if l.RemoveEnd() != nil {
		t.Error("Non-nil return value on empty List")
	}

	testReturn := func(expect int) {
		if l.RemoveEnd() != expect {
			t.Error("Wrong return value from *List.RemoveEnd")
		}
	}

	l.InsertEnd(3)
	testReturn(3)

	l.InsertEnd(3)
	l.InsertEnd(5)
	testReturn(5)
}

// Test *List.Remove
func TestList_Remove(t *testing.T) {
	l := &List{KeyCompare: func(key interface{}, obj interface{}) int8 {
		if key.(int) < obj.(int) {
			return -1
		} else if key.(int) > obj.(int) {
			return 1
		} else {
			return 0
		}
	}}

	expect := func(expect int, actual interface{}) {
		if expect != actual.(int) {
			t.Error("Expected", expect, "got", actual)
		}
	}

	l.Insert(3)
	l.Insert(5)
	l.Insert(1)
	// 1 -- 5 -- 3

	expect(1, l.Remove(1)) // 5 -- 3
	expect(5, l.head.object)

	l.Insert(1)            // 1 -- 5 -- 3
	expect(5, l.Remove(5)) // 1 -- 3
	expect(3, l.head.next.object)

	l.Insert(5)            // 5 -- 1 -- 3
	expect(3, l.Remove(3)) // 5 -- 1
	if l.tail.next != nil {
		t.Error()
	}

	if l.Remove(9) != nil {
		t.Error("Removing nonexistant val does not return nil")
	}
}
