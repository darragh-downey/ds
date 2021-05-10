package ds

import "container/heap"

// https://golang.org/pkg/container/heap/
type Item struct {
	value    string // value of the item (type is arbitrary)
	priority int    // priority in the queue
	// index required by heap.Interface methods
	index int // index of the item in the heap
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

type PQ interface {
	Insert(value int)
	Contains(value int) bool
	IsEmpty() bool
	Size() int
}

// Sedgewick; https://algs4.cs.princeton.edu/24pq/
// Heap ordered Max Priority Queue
// https://algs4.cs.princeton.edu/24pq/MaxPQ.java.html

type MaxPQ struct {
	pq []int
	qp []int
}

func NewMaxPQ(cap int) *MaxPQ {
	return &MaxPQ{
		pq: make([]int, cap),
		qp: make([]int, cap),
	}
}

func (p *MaxPQ) Insert(value int) {

}

func (p *MaxPQ) Max(value int) int {
	if p.IsEmpty() {
		return -1
	}
	// 1-based
	return p.pq[1]
}

func (p *MaxPQ) DelMax(value int) {

}

func (p *MaxPQ) IsEmpty() bool {
	return len(p.pq) == 0
}

func (p *MaxPQ) Size() int {
	return len(p.pq)
}

/***************************************************************************
 * Helper functions to restore the heap invariant.
 ***************************************************************************/

// swim - bottom-up heapify
//
//
// If the heap order is violated because a node's key becomes larger than that node's parents key,
// then we can make progress toward fixing the violation by exchanging the node with its parent.
// After the exchange, the node is larger than both its children (one is the old parent,
// and the other is smaller than the old parent because it was a child of that node) but the node may still be larger than its parent.
// We can fix that violation in the same way, and so forth, moving up the heap until we reach a node with a larger key, or the root.
func (p *MaxPQ) swim(k int) {
	for k > 1 && p.pq[k/2] < p.pq[k] {
		p.pq[k], p.pq[k/2] = p.pq[k/2], p.pq[k]
		k /= 2
	}
}

// sink - top-down heapify
//
// If the heap order is violated because a node's key becomes smaller than one or both of that node's children's keys,
// then we can make progress toward fixing the violation by exchanging the node with the larger of its two children.
// This switch may cause a violation at the child; we fix that violation in the same way, and so forth, moving down the heap
// until we reach a node with both children smaller, or the bottom.
func (p *MaxPQ) sink(k int) {
	for 2*k <= p.Size() {
		j := 2 * k
		if j < p.Size() && p.pq[j] < p.pq[j+1] {
			j++
		}
		if p.pq[k] > p.pq[j] {
			break
		}
		p.pq[k], p.pq[j] = p.pq[j], p.pq[k]
		k = j
	}
}

/***************************************************************************
 * Helper functions for compares and swaps
 ***************************************************************************/

func (p *MaxPQ) isMaxHeap() bool {
	for i := 1; i <= p.Size(); i++ {
		if p.pq[i] == 0 {
			return false
		}
	}
	// n was a variable to capture the number of pq in the priority queue
	//
	// for i := n + 1; i < p.Size(); i++ {
	// 	if p.pq[i] != 0 {
	// 		return false
	// 	}
	// }
	if p.pq[0] != 0 {
		return false
	}
	return p.isMaxHeapOrdered(1)
}

func (p *MaxPQ) isMaxHeapOrdered(k int) bool {
	if k > p.Size() {
		return true
	}
	left, right := 2*k, 2*k+1
	if left <= p.Size() && p.pq[k] < p.pq[left] {
		return false
	}
	if right <= p.Size() && p.pq[k] < p.pq[right] {
		return false
	}
	return p.isMaxHeapOrdered(left) && p.isMaxHeapOrdered(right)
}

// Sedgewick; https://algs4.cs.princeton.edu/24pq/IndexMinPQ.java.html
type MinPQ struct {
	pq []int // binary heap using 1-based indexing
	qp []int // inverse of pq - qp[pq[i]] = pq[qp[i]] = i
}

func (p *MinPQ) Insert(value int) {

}

func (p *MinPQ) Change(k int, value int) {

}

func (p *MinPQ) Contains(k int) bool {
	return false
}

func (p *MinPQ) Delete(k int) {

}

func (p *MinPQ) Min() int {
	return -1
}

func (p *MinPQ) MinIndex() int { return -1 }

// delete the minimum and return its index
func (p *MinPQ) DelMin() int {
	return -1
}

func (p *MinPQ) IsEmpty() bool {
	return len(p.pq) == 0
}

func (p *MinPQ) Size() int {
	return len(p.pq)
}
