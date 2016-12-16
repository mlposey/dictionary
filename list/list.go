package list

// ListNode is a node for a List.
type listNode struct {
	object     interface{}
	prev, next *listNode
}

// List is a double-linked list.
//
// It supports all types and does not change the order of what is inserted.
type List struct {
	head, tail *listNode
}

// Insert adds an object to the beginning of the List.
func (l *List) Insert(object interface{}) {
	node := &listNode{object: object}
	if l.head == nil {
		l.head = node
		l.tail = node
		node.next = l.head
	} else {
		l.head.prev = node
		node.next = l.head
		l.head = node
	}
}
