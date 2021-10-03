package graph_traversal

import (
	"reflect"
	"testing"
)

func TestDFSTraverse(t *testing.T) {
	alice := &Vertex{Value: "Alice"}
	cynthia := &Vertex{Value: "Cynthia"}
	bob := &Vertex{Value: "Bob"}
	alice.AdjacentVertices = append(alice.AdjacentVertices, cynthia)
	alice.AdjacentVertices = append(alice.AdjacentVertices, bob)
	cynthia.AdjacentVertices = append(cynthia.AdjacentVertices, bob)
	bob.AdjacentVertices = append(bob.AdjacentVertices, cynthia)

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
	alice.AdjacentVertices = append(alice.AdjacentVertices, cynthia)
	alice.AdjacentVertices = append(alice.AdjacentVertices, bob)
	cynthia.AdjacentVertices = append(cynthia.AdjacentVertices, bob)
	bob.AdjacentVertices = append(bob.AdjacentVertices, cynthia)

	visited := make(visitedMap)

	if !DFSSearch(alice, "Bob", visited) {
		t.Error("unexpected DFSSearch() to return")
	}
}
