package graph_traversal

import (
	"reflect"
	"testing"
)

func TestBFSTraverse(t *testing.T) {
	alice := &Vertex{Value: "Alice"}
	bob := &Vertex{Value: "Bob"}
	candy := &Vertex{Value: "Candy"}
	derek := &Vertex{Value: "Derek"}
	elaine := &Vertex{Value: "Elaine"}
	fred := &Vertex{Value: "Fred"}
	helen := &Vertex{Value: "Helen"}
	gina := &Vertex{Value: "Gina"}
	irena := &Vertex{Value: "Irena"}
	alice.AdjacentVertices = append(alice.AdjacentVertices, bob)
	alice.AdjacentVertices = append(alice.AdjacentVertices, candy)
	alice.AdjacentVertices = append(alice.AdjacentVertices, derek)
	alice.AdjacentVertices = append(alice.AdjacentVertices, elaine)
	bob.AdjacentVertices = append(alice.AdjacentVertices, fred)
	fred.AdjacentVertices = append(alice.AdjacentVertices, helen)
	candy.AdjacentVertices = append(alice.AdjacentVertices, helen)
	derek.AdjacentVertices = append(alice.AdjacentVertices, elaine)
	derek.AdjacentVertices = append(alice.AdjacentVertices, gina)
	gina.AdjacentVertices = append(alice.AdjacentVertices, irena)
	elaine.AdjacentVertices = append(alice.AdjacentVertices, derek)

	arr := BFSTraverse(alice)
	if !reflect.DeepEqual(arr, []string{"Alice", "Bob", "Candy", "Derek", "Elaine", "Fred", "Helen", "Gina", "Irena"}) {
		t.Errorf("unexpected BFSTraverse() result: %v", arr)
	}
}

func TestBFSSearch(t *testing.T) {
	alice := &Vertex{Value: "Alice"}
	bob := &Vertex{Value: "Bob"}
	candy := &Vertex{Value: "Candy"}
	derek := &Vertex{Value: "Derek"}
	elaine := &Vertex{Value: "Elaine"}
	fred := &Vertex{Value: "Fred"}
	helen := &Vertex{Value: "Helen"}
	gina := &Vertex{Value: "Gina"}
	irena := &Vertex{Value: "Irena"}
	alice.AdjacentVertices = append(alice.AdjacentVertices, bob)
	alice.AdjacentVertices = append(alice.AdjacentVertices, candy)
	alice.AdjacentVertices = append(alice.AdjacentVertices, derek)
	alice.AdjacentVertices = append(alice.AdjacentVertices, elaine)
	bob.AdjacentVertices = append(alice.AdjacentVertices, fred)
	fred.AdjacentVertices = append(alice.AdjacentVertices, helen)
	candy.AdjacentVertices = append(alice.AdjacentVertices, helen)
	derek.AdjacentVertices = append(alice.AdjacentVertices, elaine)
	derek.AdjacentVertices = append(alice.AdjacentVertices, gina)
	gina.AdjacentVertices = append(alice.AdjacentVertices, irena)
	elaine.AdjacentVertices = append(alice.AdjacentVertices, derek)

	if !BFSSearch(alice, "Irena") {
		t.Errorf("expected BFSSearch() to return true")
	}
}
