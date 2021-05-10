package ds_test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/darragh-downey/ds/go/alg"
	"github.com/darragh-downey/ds/go/ds"
)

func TestGraphDfs(t *testing.T) {
	g := ds.NewDirectedGraph()

	g.AddNode(&ds.Node{1, false})
	g.AddNode(&ds.Node{2, false})
	g.AddNode(&ds.Node{3, false})
	g.AddNode(&ds.Node{4, false})
	g.AddNode(&ds.Node{5, false})
	g.AddNode(&ds.Node{6, false})
	g.AddNode(&ds.Node{7, false})
	g.AddNode(&ds.Node{8, false})
	g.AddNode(&ds.Node{9, false})
	g.AddNode(&ds.Node{10, false})
	g.AddNode(&ds.Node{11, false})
	g.AddNode(&ds.Node{12, false})
	g.AddNode(&ds.Node{13, false})
	g.AddNode(&ds.Node{14, false})
	g.AddNode(&ds.Node{15, false})
	g.AddNode(&ds.Node{16, false})
	g.AddNode(&ds.Node{17, false})
	g.AddNode(&ds.Node{18, false})
	g.AddNode(&ds.Node{19, false})
	g.AddNode(&ds.Node{20, false})

	g.AddEdge(&ds.Node{1, false}, &ds.Node{9, false})
	g.AddEdge(&ds.Node{1, false}, &ds.Node{5, false})
	g.AddEdge(&ds.Node{1, false}, &ds.Node{2, false})
	g.AddEdge(&ds.Node{2, false}, &ds.Node{2, false})
	g.AddEdge(&ds.Node{3, false}, &ds.Node{4, false})
	g.AddEdge(&ds.Node{5, false}, &ds.Node{6, false})
	g.AddEdge(&ds.Node{5, false}, &ds.Node{8, false})
	g.AddEdge(&ds.Node{6, false}, &ds.Node{7, false})
	g.AddEdge(&ds.Node{9, false}, &ds.Node{10, false})
	g.AddEdge(&ds.Node{9, false}, &ds.Node{12, false})
	g.AddEdge(&ds.Node{10, false}, &ds.Node{13, false})
	g.AddEdge(&ds.Node{11, false}, &ds.Node{19, false})
	g.AddEdge(&ds.Node{11, false}, &ds.Node{15, false})
	g.AddEdge(&ds.Node{11, false}, &ds.Node{12, false})
	g.AddEdge(&ds.Node{12, false}, &ds.Node{12, false})
	g.AddEdge(&ds.Node{13, false}, &ds.Node{14, false})
	g.AddEdge(&ds.Node{15, false}, &ds.Node{16, false})
	g.AddEdge(&ds.Node{15, false}, &ds.Node{18, false})
	g.AddEdge(&ds.Node{16, false}, &ds.Node{17, false})
	g.AddEdge(&ds.Node{19, false}, &ds.Node{20, false})

	visitedOrder := []int{}
	cb := func(i *ds.Node) {
		visitedOrder = append(visitedOrder, int(i.Value))
	}
	alg.GraphDFS(g, &ds.Node{Value: 1, Visited: false}, make(map[int64]bool), cb)

	// add assertions here
	t.Logf("%v", visitedOrder)
}

func BenchmarkGraphDfs(b *testing.B) {
	g := ds.NewDirectedGraph()

	g.AddNode(&ds.Node{1, false})
	g.AddNode(&ds.Node{2, false})
	g.AddNode(&ds.Node{3, false})
	g.AddNode(&ds.Node{4, false})
	g.AddNode(&ds.Node{5, false})
	g.AddNode(&ds.Node{6, false})
	g.AddNode(&ds.Node{7, false})
	g.AddNode(&ds.Node{8, false})
	g.AddNode(&ds.Node{9, false})
	g.AddNode(&ds.Node{10, false})
	g.AddNode(&ds.Node{11, false})
	g.AddNode(&ds.Node{12, false})
	g.AddNode(&ds.Node{13, false})
	g.AddNode(&ds.Node{14, false})
	g.AddNode(&ds.Node{15, false})
	g.AddNode(&ds.Node{16, false})
	g.AddNode(&ds.Node{17, false})
	g.AddNode(&ds.Node{18, false})
	g.AddNode(&ds.Node{19, false})
	g.AddNode(&ds.Node{20, false})

	g.AddEdge(&ds.Node{1, false}, &ds.Node{9, false})
	g.AddEdge(&ds.Node{1, false}, &ds.Node{5, false})
	g.AddEdge(&ds.Node{1, false}, &ds.Node{2, false})
	g.AddEdge(&ds.Node{2, false}, &ds.Node{2, false})
	g.AddEdge(&ds.Node{3, false}, &ds.Node{4, false})
	g.AddEdge(&ds.Node{5, false}, &ds.Node{6, false})
	g.AddEdge(&ds.Node{5, false}, &ds.Node{8, false})
	g.AddEdge(&ds.Node{6, false}, &ds.Node{7, false})
	g.AddEdge(&ds.Node{9, false}, &ds.Node{10, false})
	g.AddEdge(&ds.Node{9, false}, &ds.Node{12, false})
	g.AddEdge(&ds.Node{10, false}, &ds.Node{13, false})
	g.AddEdge(&ds.Node{11, false}, &ds.Node{19, false})
	g.AddEdge(&ds.Node{11, false}, &ds.Node{15, false})
	g.AddEdge(&ds.Node{11, false}, &ds.Node{12, false})
	g.AddEdge(&ds.Node{12, false}, &ds.Node{12, false})
	g.AddEdge(&ds.Node{13, false}, &ds.Node{14, false})
	g.AddEdge(&ds.Node{15, false}, &ds.Node{16, false})
	g.AddEdge(&ds.Node{15, false}, &ds.Node{18, false})
	g.AddEdge(&ds.Node{16, false}, &ds.Node{17, false})
	g.AddEdge(&ds.Node{19, false}, &ds.Node{20, false})

	visitedOrder := []int{}
	cb := func(i *ds.Node) {
		visitedOrder = append(visitedOrder, int(i.Value))
	}
	for i := 0; i < b.N; i++ {
		alg.GraphDFS(g, &ds.Node{Value: 1, Visited: false}, make(map[int64]bool), cb)
	}
}

func TestGraphDfsCo(t *testing.T) {
	g := ds.NewDirectedGraph()

	g.AddNode(&ds.Node{1, false})
	g.AddNode(&ds.Node{2, false})
	g.AddNode(&ds.Node{3, false})
	g.AddNode(&ds.Node{4, false})
	g.AddNode(&ds.Node{5, false})
	g.AddNode(&ds.Node{6, false})
	g.AddNode(&ds.Node{7, false})
	g.AddNode(&ds.Node{8, false})
	g.AddNode(&ds.Node{9, false})
	g.AddNode(&ds.Node{10, false})
	g.AddNode(&ds.Node{11, false})
	g.AddNode(&ds.Node{12, false})
	g.AddNode(&ds.Node{13, false})
	g.AddNode(&ds.Node{14, false})
	g.AddNode(&ds.Node{15, false})
	g.AddNode(&ds.Node{16, false})
	g.AddNode(&ds.Node{17, false})
	g.AddNode(&ds.Node{18, false})
	g.AddNode(&ds.Node{19, false})
	g.AddNode(&ds.Node{20, false})

	g.AddEdge(&ds.Node{1, false}, &ds.Node{9, false})
	g.AddEdge(&ds.Node{1, false}, &ds.Node{5, false})
	g.AddEdge(&ds.Node{1, false}, &ds.Node{2, false})
	g.AddEdge(&ds.Node{2, false}, &ds.Node{2, false})
	g.AddEdge(&ds.Node{3, false}, &ds.Node{4, false})
	g.AddEdge(&ds.Node{5, false}, &ds.Node{6, false})
	g.AddEdge(&ds.Node{5, false}, &ds.Node{8, false})
	g.AddEdge(&ds.Node{6, false}, &ds.Node{7, false})
	g.AddEdge(&ds.Node{9, false}, &ds.Node{10, false})
	g.AddEdge(&ds.Node{9, false}, &ds.Node{12, false})
	g.AddEdge(&ds.Node{10, false}, &ds.Node{13, false})
	g.AddEdge(&ds.Node{11, false}, &ds.Node{19, false})
	g.AddEdge(&ds.Node{11, false}, &ds.Node{15, false})
	g.AddEdge(&ds.Node{11, false}, &ds.Node{12, false})
	g.AddEdge(&ds.Node{12, false}, &ds.Node{12, false})
	g.AddEdge(&ds.Node{13, false}, &ds.Node{14, false})
	g.AddEdge(&ds.Node{15, false}, &ds.Node{16, false})
	g.AddEdge(&ds.Node{15, false}, &ds.Node{18, false})
	g.AddEdge(&ds.Node{16, false}, &ds.Node{17, false})
	g.AddEdge(&ds.Node{19, false}, &ds.Node{20, false})

	wg := &sync.WaitGroup{}
	visitedOrder := []int{}
	cb := func(i *ds.Node) {
		visitedOrder = append(visitedOrder, int(i.Value))
	}
	wg.Add(1)
	alg.GraphDFSCo(g, &ds.Node{Value: 1, Visited: false}, &sync.Map{}, wg, cb)

	// add assertions here
	t.Logf("%v", visitedOrder)
}

func BenchmarkGraphDfsCo(b *testing.B) {
	g := ds.NewDirectedGraph()

	g.AddNode(&ds.Node{1, false})
	g.AddNode(&ds.Node{2, false})
	g.AddNode(&ds.Node{3, false})
	g.AddNode(&ds.Node{4, false})
	g.AddNode(&ds.Node{5, false})
	g.AddNode(&ds.Node{6, false})
	g.AddNode(&ds.Node{7, false})
	g.AddNode(&ds.Node{8, false})
	g.AddNode(&ds.Node{9, false})
	g.AddNode(&ds.Node{10, false})
	g.AddNode(&ds.Node{11, false})
	g.AddNode(&ds.Node{12, false})
	g.AddNode(&ds.Node{13, false})
	g.AddNode(&ds.Node{14, false})
	g.AddNode(&ds.Node{15, false})
	g.AddNode(&ds.Node{16, false})
	g.AddNode(&ds.Node{17, false})
	g.AddNode(&ds.Node{18, false})
	g.AddNode(&ds.Node{19, false})
	g.AddNode(&ds.Node{20, false})

	g.AddEdge(&ds.Node{1, false}, &ds.Node{9, false})
	g.AddEdge(&ds.Node{1, false}, &ds.Node{5, false})
	g.AddEdge(&ds.Node{1, false}, &ds.Node{2, false})
	g.AddEdge(&ds.Node{2, false}, &ds.Node{2, false})
	g.AddEdge(&ds.Node{3, false}, &ds.Node{4, false})
	g.AddEdge(&ds.Node{5, false}, &ds.Node{6, false})
	g.AddEdge(&ds.Node{5, false}, &ds.Node{8, false})
	g.AddEdge(&ds.Node{6, false}, &ds.Node{7, false})
	g.AddEdge(&ds.Node{9, false}, &ds.Node{10, false})
	g.AddEdge(&ds.Node{9, false}, &ds.Node{12, false})
	g.AddEdge(&ds.Node{10, false}, &ds.Node{13, false})
	g.AddEdge(&ds.Node{11, false}, &ds.Node{19, false})
	g.AddEdge(&ds.Node{11, false}, &ds.Node{15, false})
	g.AddEdge(&ds.Node{11, false}, &ds.Node{12, false})
	g.AddEdge(&ds.Node{12, false}, &ds.Node{12, false})
	g.AddEdge(&ds.Node{13, false}, &ds.Node{14, false})
	g.AddEdge(&ds.Node{15, false}, &ds.Node{16, false})
	g.AddEdge(&ds.Node{15, false}, &ds.Node{18, false})
	g.AddEdge(&ds.Node{16, false}, &ds.Node{17, false})
	g.AddEdge(&ds.Node{19, false}, &ds.Node{20, false})

	wg := &sync.WaitGroup{}
	visitedOrder := []int{}
	cb := func(i *ds.Node) {
		visitedOrder = append(visitedOrder, int(i.Value))
	}
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		alg.GraphDFSCo(g, &ds.Node{Value: 1, Visited: false}, &sync.Map{}, wg, cb)
	}
}

func TestBfs(t *testing.T) {
	g := ds.NewDirectedGraph()

	g.AddNode(&ds.Node{1, false})
	g.AddNode(&ds.Node{2, false})
	g.AddNode(&ds.Node{3, false})
	g.AddNode(&ds.Node{4, false})
	g.AddNode(&ds.Node{5, false})
	g.AddNode(&ds.Node{6, false})
	g.AddNode(&ds.Node{7, false})
	g.AddNode(&ds.Node{8, false})
	g.AddNode(&ds.Node{9, false})
	g.AddNode(&ds.Node{10, false})

	g.AddEdge(&ds.Node{1, false}, &ds.Node{9, false})
	g.AddEdge(&ds.Node{1, false}, &ds.Node{5, false})
	g.AddEdge(&ds.Node{1, false}, &ds.Node{2, false})
	g.AddEdge(&ds.Node{2, false}, &ds.Node{2, false})
	g.AddEdge(&ds.Node{3, false}, &ds.Node{4, false})
	g.AddEdge(&ds.Node{5, false}, &ds.Node{6, false})
	g.AddEdge(&ds.Node{5, false}, &ds.Node{8, false})
	g.AddEdge(&ds.Node{6, false}, &ds.Node{7, false})
	g.AddEdge(&ds.Node{9, false}, &ds.Node{10, false})

	visitedOrder := []int{}
	cb := func(i *ds.Node) {
		visitedOrder = append(visitedOrder, int(i.Value))
	}
	alg.BFS(g, 1, cb)

	// add assertions here
	fmt.Println(visitedOrder)
}
