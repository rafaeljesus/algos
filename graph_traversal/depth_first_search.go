package graph_traversal

type Vertex struct {
	Value            string
	AdjacentVertices []*Vertex
}

type visitedMap map[string]struct{}

func DFSTraverse(v *Vertex, visited visitedMap, arr []string) []string {
	arr = append(arr, v.Value)
	visited[v.Value] = struct{}{}
	for _, vertex := range v.AdjacentVertices {
		if _, ok := visited[vertex.Value]; ok {
			continue
		}
		arr = DFSTraverse(vertex, visited, arr)
	}
	return arr
}

func DFSSearch(v *Vertex, search string, visited visitedMap) bool {
	if v.Value == search {
		return true
	}
	visited[v.Value] = struct{}{}
	for _, vertex := range v.AdjacentVertices {
		if _, ok := visited[vertex.Value]; ok {
			continue
		}
		return DFSSearch(vertex, search, visited)
	}
	return false
}
