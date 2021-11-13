package graph_traversal

func BFSTraverse(v *Vertex) []string {
	visited := make(map[string]struct{})
	visited[v.Value] = struct{}{}
	arr := []string{}
	arr = append(arr, v.Value)
	queue := v.AdjacentVertices
	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]
		if _, ok := visited[vertex.Value]; !ok {
			queue = append(queue, vertex.AdjacentVertices...)
			arr = append(arr, vertex.Value)
			visited[vertex.Value] = struct{}{}
		}
	}
	return arr
}

func BFSSearch(v *Vertex, searchTerm string) bool {
	if v.Value == searchTerm {
		return true
	}
	visited := make(map[string]struct{})
	queue := v.AdjacentVertices
	for len(queue) > 0 {
		vertex := queue[0]
		if vertex.Value == searchTerm {
			return true
		}
		queue = queue[1:]
		if _, ok := visited[vertex.Value]; !ok {
			visited[v.Value] = struct{}{}
			queue = append(queue, vertex.AdjacentVertices...)
		}
	}
	return false
}
