package list

// ListNode is a node for a List.
type listNode struct {
	object     interface{}
	prev, next *listNode
}

// List is a doubly-linked list.
//
// It supports all types and does not change the order of what is inserted.
//
// KeyCompare -
//	This required function should compare a key to an object which
//	contains a key of that type. The return values should be:
//		-1 if key is less than object key
//		0 if key is equal to object key
//		1 if key is greater than object key
type List struct {
	KeyCompare func(interface{}, interface{}) int8
	Size       uint
	head, tail *listNode
}

// Insert adds an object to the beginning of the List.
func (l *List) Insert(object interface{}) {
	node := &listNode{object: object}
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.head.prev = node
		node.next = l.head
		l.head = node
	}
	l.Size++
}

// InsertEnd adds an object to the end of the List.
func (l *List) InsertEnd(object interface{}) {
	node := &listNode{object: object}
	if l.head == nil {
		l.head = node
		l.tail = node
		l.tail.prev = l.head
	} else {
		l.tail.next = node
		node.prev = l.tail
		l.tail = node
	}
	l.Size++
}

// RemoveFront removes and returns the front item.
//
// Returns nil if the *List is empty.
func (l *List) RemoveFront() interface{} {
	if l.head == nil {
		return nil
	}
	node := l.head
	l.head = l.head.next
	l.Size--
	return node.object
}

// RemoveEnd removes the last item.
//
// Returns nil if the *List is empty.
func (l *List) RemoveEnd() interface{} {
	if l.head == nil {
		return nil
	}
	node := l.tail
	l.tail = l.tail.prev
	l.Size--
	return node.object
}

// Remove removes and returns an object (with a matching key) from the List.
//
// The comparison is done with List.KeyCompare.
func (l *List) Remove(key interface{}) interface{} {
	n := l.head
	for n != nil {
		if l.KeyCompare(key, n.object) == 0 {
			if n.prev != nil {
				n.prev.next = n.next
			}
			if n.next != nil {
				n.next.prev = n.prev
			}
			if n == l.head {
				l.head = l.head.next
			}
			if n == l.tail {
				l.tail = l.tail.prev
			}
			l.Size--
			return n.object
		}
		n = n.next
	}
	return nil
}
