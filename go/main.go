package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/darragh-downey/ds/go/ds"
)

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
	processors := runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Printf("\nTime elapsed: %v\n\n", time.Since(time.Now()))

	wg := sync.WaitGroup{}
	wg.Add(1)
	go b.DFSParallel(&wg)
	wg.Wait()

	fmt.Printf("\nProcessors: %d\n\n", processors)
}
