package list

// Object defines an interface for List values.
//
// All inserted values should implement the operations defined here.
type Object interface {
	// Return true if other is less than this.
	Less(other Object) bool
}

// ListNode is a node for a List.
type ListNode struct {
	object      *Object
	left, right *ListNode
}

// List is a double-linked list.
// TODO: Improve List documentation.
type List struct {
	head, tail *ListNode
}
