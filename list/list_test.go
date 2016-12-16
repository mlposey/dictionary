package list

import "testing"

// Test List.Insert
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
