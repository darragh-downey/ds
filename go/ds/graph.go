package ds

import (
	"fmt"
	"sync"
)

type Node struct {
	Value int64
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.Value)
}

type Graph struct {
	Nodes []*Node
	Edges map[int64][]*Node
}

func (g *Graph) AddNode(n *Node) {
	g.Nodes = append(g.Nodes, n)
}

// AddEdge Unidirectional
func (g *Graph) AddEdge(n1, n2 *Node) {
	if g.Edges == nil {
		g.Edges = make(map[int64][]*Node)
	}

	g.Edges[n1.Value] = append(g.Edges[n1.Value], n2)
}

func (g *Graph) AddEdgeBi(n1, n2 *Node) {
	g.AddEdge(n1, n2)
	g.AddEdge(n2, n1)
}

func (g *Graph) ExistingNode(n *Node) Node {
	for i := 0; i < len(g.Nodes); i++ {
		if n.Value == g.Nodes[i].Value {
			return *g.Nodes[i]
		}
	}

	g.AddNode(n)
	return *n
}

func (g *Graph) String() {
	s := ""
	for i := 0; i < len(g.Nodes); i++ {
		s += g.Nodes[i].String() + "-> "
		near := g.Edges[g.Nodes[i].Value]
		for j := 0; j < len(near); j++ {
			s += near[j].String() + " "
		}
		s += "\n"
	}
	fmt.Println(s)
}

type NodeQueue struct {
	items []Node
	lock  sync.RWMutex
}

func (pq NodeQueue) New() *NodeQueue {
	return &NodeQueue{
		[]Node{},
		sync.RWMutex{},
	}
}

func (pq *NodeQueue) Enqueue(n Node) {
	pq.lock.Lock()
	defer pq.lock.Unlock()
	pq.items = append(pq.items, n)
}

func (pq *NodeQueue) Dequeue() *Node {
	pq.lock.Lock()
	defer pq.lock.Unlock()
	n := pq.items[0]
	pq.items = pq.items[1:len(pq.items)]
	return &n
}

func (pq *NodeQueue) Front() *Node {
	pq.lock.Lock()
	defer pq.lock.Unlock()
	return &pq.items[0]
}

func (pq *NodeQueue) IsEmpty() bool {
	pq.lock.Lock()
	defer pq.lock.Unlock()
	return len(pq.items) == 0
}

func (pq *NodeQueue) Size() int {
	pq.lock.Lock()
	defer pq.lock.Unlock()
	return len(pq.items)
}

// BFS Breadth First Search
func (g *Graph) BFS(idx int, f func(*Node)) (connections int64) {
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

// DFS Depth First Search
func (g *Graph) DFS() {}

func (g *Graph) DFSParallel(wg *sync.WaitGroup) {
	defer wg.Done()
}
