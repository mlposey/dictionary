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

// TODO: This should be *List.InsertEnd.
// Test List.InsertEnd
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

	if l.RemoveFront() != nil {
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
