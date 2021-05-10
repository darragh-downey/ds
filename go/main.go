package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/darragh-downey/ds/go/alg"
	"github.com/darragh-downey/ds/go/ds"
)

func main() {
	b := &ds.BST{}

	b.Insert(50)
	b.Insert(30)
	b.Insert(20)
	b.Insert(40)
	b.Insert(70)
	b.Insert(60)
	b.Insert(80)

	b.Inorder()
	processors := runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Printf("\nTime elapsed: %v\n\n", time.Since(time.Now()))

	wg := sync.WaitGroup{}
	wg.Add(1)
	go alg.BSTDFSCo(b, &wg)
	wg.Wait()

	fmt.Printf("\nProcessors: %d\n\n", processors)
}
