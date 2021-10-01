package depth_first_search

import (
	"reflect"
	"testing"
)

func TestDFSTraverse(t *testing.T) {
	alice := &Vertex{Value: "Alice"}
	cynthia := &Vertex{Value: "Cynthia"}
	bob := &Vertex{Value: "Bob"}
	alice.AdjacentVertex = append(alice.AdjacentVertex, cynthia)
	alice.AdjacentVertex = append(alice.AdjacentVertex, bob)
	cynthia.AdjacentVertex = append(cynthia.AdjacentVertex, bob)
	bob.AdjacentVertex = append(bob.AdjacentVertex, cynthia)

	visited := make(visitedMap)

	arr := DFSTraverse(alice, visited, []string{})
	if !reflect.DeepEqual(arr, []string{"Alice", "Cynthia", "Bob"}) {
		t.Errorf("unexpected DFSTraverse() result: %v", arr)
	}
}

func TestDFSSearch(t *testing.T) {
	alice := &Vertex{Value: "Alice"}
	cynthia := &Vertex{Value: "Cynthia"}
	bob := &Vertex{Value: "Bob"}
	alice.AdjacentVertex = append(alice.AdjacentVertex, cynthia)
	alice.AdjacentVertex = append(alice.AdjacentVertex, bob)
	cynthia.AdjacentVertex = append(cynthia.AdjacentVertex, bob)
	bob.AdjacentVertex = append(bob.AdjacentVertex, cynthia)

	visited := make(visitedMap)

	if !DFSSearch(alice, "Bob", visited) {
		t.Error("unexpected DFSSearch() to return")
	}
}
