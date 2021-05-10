package alg

import "sync"

// DFS Depth First Search
func (g *ds.Graph) DFS(root *ds.Node, visit func(*Node)) {
	if root == nil {
		return
	}
	visit(root)
	root.visited = true
	for _, node := range g.Edges[node] {
		if visited[node] == false {
			g.DFS(*node)
		}
	}
}

func DFSParallel(g *ds.Graph, idx int, wg *sync.WaitGroup, f *ds.Node) {
	defer wg.Done()
}
