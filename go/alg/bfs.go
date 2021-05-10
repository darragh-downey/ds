package alg

import (
	"sync"

	"container/list"

	"github.com/darragh-downey/ds/go/ds"
)

// BFS Breadth First Search
func BFS(g *ds.Graph, idx int, f func(*ds.Node)) (connections int64) {
	q := list.New()

	n := g.Nodes[idx]
	q.PushBack(n)
	visited := make(map[int32]bool)

	for q.Len() > 0 {
		node := q.Front().Value
		visited[int32(node.(ds.Node).Value)] = true
		near := g.Edges[node.(ds.Node).Value]

		for i := 0; i < len(near); i++ {
			j := near[i]
			if !visited[int32(j.Value)] {
				q.PushBack(j)
				visited[int32(j.Value)] = true
			}
		}
		if f != nil {
			f(node.(*ds.Node))
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

func GraphBFSCo(g *ds.Graph, idx int, q list.List, discovered *sync.Map, wg *sync.WaitGroup, f func(*ds.Node)) {

}
