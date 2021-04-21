package main

import "github.com/robatussum/ds/go/ds"

func main() {
	b, root := &ds.BST{}, &ds.BST{}

	root = b.Insert(root, 50)
	b.Insert(root, 30)
	b.Insert(root, 20)
	b.Insert(root, 40)
	b.Insert(root, 70)
	b.Insert(root, 60)
	b.Insert(root, 80)

	b.Inorder(root)
}
