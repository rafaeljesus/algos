package graph_traversal

func BFSTraverse(v *Vertex) []string {
	visited := make(visitedMap)
	visited[v.Value] = struct{}{}
	arr := make([]string, 0)
	arr = append(arr, v.Value)
	queue := make([]*Vertex, 0)
	queue = append(queue, v)
	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]
		for _, av := range vertex.AdjacentVertices {
			if _, ok := visited[av.Value]; ok {
				continue
			}
			visited[av.Value] = struct{}{}
			queue = append(queue, av)
			arr = append(arr, av.Value)
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
