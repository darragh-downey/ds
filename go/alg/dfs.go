package alg

import (
	"fmt"
	"sync"

	"github.com/darragh-downey/ds/go/ds"
)

// GraphDFS
// An implementation of Depth First Search on a graph
func GraphDFS(g *ds.Graph, root *ds.Node, visited map[int64]bool, visit func(*ds.Node)) {
	if root == nil {
		return
	}
	visit(root)
	visited[root.Value] = true
	root.Visited = true
	for _, node := range g.Edges[root.Value] {
		if !visited[node.Value] {
			GraphDFS(g, node, visited, visit)
		}
	}
}

// GraphDFSCo
// A goroutine implementation of Depth First Search on a graph
// Makes use of the sync.Map struct for concurrent R/W
func GraphDFSCo(g *ds.Graph, root *ds.Node, visited *sync.Map, wg *sync.WaitGroup, visit func(*ds.Node)) {
	defer wg.Done()
	lock := sync.RWMutex{}
	lock.Lock()
	defer lock.Unlock()

	if root == nil {
		return
	}

	visit(root)

	visited.Store(root.Value, true)
	root.Visited = true
	for _, node := range g.Edges[root.Value] {
		if _, ok := visited.Load(node.Value); !ok {
			wg.Add(1)
			go GraphDFSCo(g, node, visited, wg, visit)
		}
	}

}

// BSTDFS
// An implementation of Binary Search Tree Depth First Search
func BSTDFS(b *ds.BST) {
	if b == nil {
		return
	}

	BSTDFS(b.Left)
	fmt.Printf("%v", b.Value)
	BSTDFS(b.Right)
}

// BSTDFSCo
// A goroutine implementation of a Binary Search Tree Depth First Search
func BSTDFSCo(b *ds.BST, wg *sync.WaitGroup) {
	defer wg.Done()

	if b == nil {
		return
	}

	wg.Add(1)
	go BSTDFSCo(b.Left, wg)

	fmt.Printf("%v", b)

	wg.Add(1)
	go BSTDFSCo(b.Right, wg)
}
