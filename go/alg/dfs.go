package alg

import (
	"sync"

	"github.com/darragh-downey/ds/go/ds"
)

// DFS Depth First Search
func DFS(g *ds.Graph, root *ds.Node, visited map[int64]bool, visit func(*ds.Node)) {
	if root == nil {
		return
	}
	visit(root)
	root.Visited = true
	for _, node := range g.Edges[root.Value] {
		if !visited[node.Value] {
			DFS(g, node, visited, visit)
		}
	}
}

func DFSParallel(g *ds.Graph, idx int, wg *sync.WaitGroup, f *ds.Node) {
	defer wg.Done()
}
