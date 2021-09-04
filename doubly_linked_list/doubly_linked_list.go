package doubly_linked_list

type Node struct {
	next, prev *Node
	Value      interface{}
}

func (ll *Node) Next() *Node {
	return ll.next
}

type List struct {
	Head, Tail *Node
}

// 1 <-> 2 <-> 3 <-> 4 <-> 5
// ll.InsertToFront(7)
// 7 <-> 1 <-> 2 <-> 3 <-> 4 <-> 5
func (ll *List) InsertToFront(n *Node) {
	if ll.Head == nil {
		ll.Head = n
		ll.Tail = n
		return
	}
	n.next = ll.Head
	ll.Head.next = n
	ll.Head = n
}

// 7 <-> 1 <-> 2 <-> 3 <-> 4 <-> 5
// ll.InsertToBack(6)
// 7 <-> 1 <-> 2 <-> 3 <-> 4 <-> 5 <> 6
func (ll *List) InsertToBack(n *Node) {
	if ll.Tail == nil {
		ll.InsertToFront(n)
		return
	}
	n.prev = ll.Tail
	ll.Tail.next = n
	ll.Tail = n
}

// 7 <-> 1 <-> 2 <-> 3 <-> 4 <-> 5 <> 6
// ll.MoveToFront(3)
// 3 <-> 7 <-> 1 <-> 2 <-> 4 <-> 5 <> 6
func (ll *List) MoveToFront(n *Node) {
	if ll.Head == n {
		return
	}
	ll.Remove(n)
	n.prev = ll.Head.prev
	n.next = ll.Head
	ll.Head = n
	ll.Head.prev = n
}

// 7 <-> 1 <-> 2 <-> 3 <-> 4 <-> 5 <> 6
// ll.MoveToBack(3)
// 7 <-> 1 <-> 2 <-> 4 <-> 5 <> 6 <> 3
func (ll *List) MoveToBack(n *Node) {
	if ll.Tail == n {
		return
	}
	ll.Remove(n)
	n.prev = ll.Tail
	n.next = ll.Tail.next
	ll.Tail = n
	ll.Tail.next = n
}

// 7 <-> 1 <-> 2 <-> 3 <-> 4 <-> 5 <> 6
// ll.Remove(3)
// 7 <-> 1 <-> 2 <-> 4 <-> 5 <> 6
func (ll *List) Remove(n *Node) {
	if ll.Head == n {
		ll.Head = ll.Head.next
	}
	if ll.Tail == n {
		ll.Tail = ll.Tail.next
	}
	if n.prev != nil {
		n.prev.next = n.next
	}
	if n.next != nil {
		n.next.prev = n.prev
	}
}
