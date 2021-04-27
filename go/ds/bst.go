package ds

import (
	"fmt"
	"sync"
)

type BST struct {
	Value int
	Root  bool
	Left  *BST
	Right *BST
}

// Search recursive search
func (b *BST) Search(key int) *BST {
	if b == nil || b.Value == key {
		return b
	}

	if b.Value < key {
		return b.Right.Search(key)
	}

	return b.Left.Search(key)
}

func (b *BST) Insert(value int) *BST {
	if b == nil {
		b.Value = value
		return &BST{
			Value: value,
		}
	}

	if value > b.Value {
		b.Right = b.Right.Insert(value)
	} else {
		b.Left = b.Left.Insert(value)
	}

	return b
}

// Inorder traversal
// left, root, right
func (b *BST) Inorder() {
	if b == nil {
		return
	}

	b.Left.Inorder()
	fmt.Printf("%v", b.Value)
	b.Right.Inorder()
}

// Preorder traversal
// root, left, right
func (b *BST) Preorder() {
	if b == nil {
		return
	}

	fmt.Printf("%v", b.Value)
	b.Left.Preorder()
	b.Right.Preorder()
}

// Postorder traversal
// left, right, root
func (b *BST) Postorder() {
	if b == nil {
		return
	}

	b.Left.Postorder()
	b.Right.Postorder()
	fmt.Printf("%v", b.Value)
}

func (b *BST) DFS() {
	if b == nil {
		return
	}

	b.Left.DFS()
	fmt.Printf("%v", b.Value)
	b.Right.DFS()
}

func (b *BST) ProcessNodeParallel(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%v", b.Value)
}

func (b *BST) DFSParallel(wg *sync.WaitGroup) {
	defer wg.Done()

	if b == nil {
		return
	}

	wg.Add(1)
	go b.Left.DFSParallel(wg)

	wg.Add(1)
	go b.ProcessNodeParallel(wg)

	wg.Add(1)
	go b.Right.DFSParallel(wg)
}
