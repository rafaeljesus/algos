package doubly_linked_list

import (
	"testing"
)

func Test(t *testing.T) {
	bob := &Node{Value: "Bob"}
	jill := &Node{Value: "Jill"}
	emily := &Node{Value: "Emily"}
	greg := &Node{Value: "Greg"}
	ll := &List{}
	ll.InsertToFront(bob)
	ll.InsertToFront(jill)
	ll.InsertToFront(emily)
	ll.InsertToFront(greg)
	if ll.Head.Value.(string) != "Greg" {
		t.Fatal("ll.Head should be Greg")
	}
	ll.MoveToFront(jill)
	if ll.Head.Value.(string) != "Jill" {
		t.Fatal("ll.Head should be Jill")
	}
	ll.InsertToBack(jill)
	if ll.Tail.Value.(string) != "Jill" {
		t.Fatal("ll.Tail should be Jill")
	}
	ll.MoveToBack(greg)
	if ll.Tail.Value.(string) != "Greg" {
		t.Fatal("ll.Tail should be Greg")
	}
}
