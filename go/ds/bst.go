package ds

import "fmt"

type BST struct {
	Value int
	Root  bool
	Left  *BST
	Right *BST
}

func (b *BST) Search(root *BST, key int) *BST {
	if b == nil || b.Value == key {
		return b
	}

	if b.Value < key {
		return b.Search(b.Right, key)
	}

	return b.Search(b.Left, key)
}

func (b *BST) Insert(root *BST, value int) *BST {
	if b == nil {
		b.Value = value
		return &BST{
			Value: value,
		}
	}

	if value > b.Value {
		b.Right = b.Insert(b.Right, value)
	} else {
		b.Left = b.Insert(b.Left, value)
	}

	return root
}

func (b *BST) Inorder(root *BST) {
	if b == nil {
		return
	}

	b.Inorder(root.Left)
	fmt.Printf("%v", b.Value)
	b.Inorder(root.Right)
}
