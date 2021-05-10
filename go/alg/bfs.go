package alg

// BFS Breadth First Search
func BFS(g *Graph, idx int, f func(*Node)) (connections int64) {
	q := NodeQueue{}
	q.New()

	n := g.Nodes[idx]
	q.Enqueue(*n)
	visited := make(map[int32]bool)

	for {
		if q.IsEmpty() {
			break
		}

		node := q.Dequeue()
		visited[int32(node.Value)] = true
		near := g.Edges[node.Value]

		for i := 0; i < len(near); i++ {
			j := near[i]
			if !visited[int32(j.Value)] {
				q.Enqueue(*j)
				visited[int32(j.Value)] = true
			}
		}
		if f != nil {
			f(node)
		}
	}

	connections = int64(0)
	for _, v := range visited {
		if v {
			connections++
		}
	}
	return
}
