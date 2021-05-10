package alg_test

import (
	"testing"

	"github.com/darragh-downey/ds/go/alg"
)

func TestBUMergesort(t *testing.T) {
	var cases = []struct {
		i []int
	}{
		{},
		{},
		{},
	}

	for _, c := range cases {
		alg.BottomUpMerge(c.i)

		if !isSorted(c.i) {
			t.Errorf("E: Not sorted %v", c.i)
		}
	}
}

func isSorted(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}
