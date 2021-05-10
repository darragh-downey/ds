package ds

import (
	"fmt"
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
