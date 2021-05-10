package test

import (
	"fmt"
	"testing"

	"github.com/darragh-downey/ds/go/alg"
	"github.com/darragh-downey/ds/go/ds"
)

func TestDfs(t *testing.T) {
	g := ds.NewDirectedGraph()

	g.AddNode(&ds.Node{1})
	g.AddNode(&ds.Node{2})
	g.AddNode(&ds.Node{3})
	g.AddNode(&ds.Node{4})
	g.AddNode(&ds.Node{5})
	g.AddNode(&ds.Node{6})
	g.AddNode(&ds.Node{7})
	g.AddNode(&ds.Node{8})
	g.AddNode(&ds.Node{9})
	g.AddNode(&ds.Node{10})

	g.AddEdge(&ds.Node{1}, &ds.Node{9})
	g.AddEdge(&ds.Node{1}, &ds.Node{5})
	g.AddEdge(&ds.Node{1}, &ds.Node{2})
	g.AddEdge(&ds.Node{2}, &ds.Node{2})
	g.AddEdge(&ds.Node{3}, &ds.Node{4})
	g.AddEdge(&ds.Node{5}, &ds.Node{6})
	g.AddEdge(&ds.Node{5}, &ds.Node{8})
	g.AddEdge(&ds.Node{6}, &ds.Node{7})
	g.AddEdge(&ds.Node{9}, &ds.Node{10})

	visitedOrder := []int{}
	cb := func(i int) {
		visitedOrder = append(visitedOrder, i)
	}
	alg.DFS(g, g.Nodes[&ds.Node{1}], cb)

	// add assertions here
	fmt.Println(visitedOrder)
}

func TestBfs(t *testing.T) {
	g := ds.NewDirectedGraph()

	g.AddNode(&ds.Node{1})
	g.AddNode(&ds.Node{2})
	g.AddNode(&ds.Node{3})
	g.AddNode(&ds.Node{4})
	g.AddNode(&ds.Node{5})
	g.AddNode(&ds.Node{6})
	g.AddNode(&ds.Node{7})
	g.AddNode(&ds.Node{8})
	g.AddNode(&ds.Node{9})
	g.AddNode(&ds.Node{10})

	g.AddEdge(&ds.Node{1}, &ds.Node{9})
	g.AddEdge(&ds.Node{1}, &ds.Node{5})
	g.AddEdge(&ds.Node{1}, &ds.Node{2})
	g.AddEdge(&ds.Node{2}, &ds.Node{2})
	g.AddEdge(&ds.Node{3}, &ds.Node{4})
	g.AddEdge(&ds.Node{5}, &ds.Node{6})
	g.AddEdge(&ds.Node{5}, &ds.Node{8})
	g.AddEdge(&ds.Node{6}, &ds.Node{7})
	g.AddEdge(&ds.Node{9}, &ds.Node{10})

	visitedOrder := []int{}
	cb := func(i int) {
		visitedOrder = append(visitedOrder, i)
	}
	alg.BFS(g, g.Nodes[1], cb)

	// add assertions here
	fmt.Println(visitedOrder)
}
