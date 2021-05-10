package ds_test

import (
	"testing"

	"github.com/darragh-downey/ds/go/ds"
)

func TestBST(t *testing.T) {
	var cases = []struct {
		expected bool
	}{}

	for _, c := range cases {
		if !c.expected {
			t.Errorf("E: Placeholder")
		}
	}
}

func TestBSTDFS(t *testing.T) {
	var cases = []struct {
		bst      *ds.BST
		expected bool
	}{
		{
			&ds.BST{
				Value: 1,
				Root:  true,
				Left: &ds.BST{
					Value: 2,
					Root:  true,
					Left: &ds.BST{
						Value: 4,
						Root:  true,
						Left:  &ds.BST{},
						Right: &ds.BST{},
					},
					Right: &ds.BST{
						Value: 5,
						Root:  true,
						Left:  &ds.BST{},
						Right: &ds.BST{},
					},
				},
				Right: &ds.BST{
					Value: 3,
					Root:  true,
					Left: &ds.BST{
						Value: 6,
						Root:  true,
						Left:  &ds.BST{},
						Right: &ds.BST{},
					},
					Right: &ds.BST{
						Value: 7,
						Root:  true,
						Left:  &ds.BST{},
						Right: &ds.BST{},
					},
				},
			},
			true,
		},
	}

	for _, c := range cases {
		if !c.expected {
			t.Errorf("E: Placeholder")
		}
	}
}

func TestBSTDFSParallel(t *testing.T) {
	var cases = []struct {
		bst      *ds.BST
		expected bool
	}{
		{
			&ds.BST{
				Value: 1,
				Root:  true,
				Left: &ds.BST{
					Value: 2,
					Root:  true,
					Left: &ds.BST{
						Value: 4,
						Root:  true,
						Left:  &ds.BST{},
						Right: &ds.BST{},
					},
					Right: &ds.BST{
						Value: 5,
						Root:  true,
						Left:  &ds.BST{},
						Right: &ds.BST{},
					},
				},
				Right: &ds.BST{
					Value: 3,
					Root:  true,
					Left: &ds.BST{
						Value: 6,
						Root:  true,
						Left:  &ds.BST{},
						Right: &ds.BST{},
					},
					Right: &ds.BST{
						Value: 7,
						Root:  true,
						Left:  &ds.BST{},
						Right: &ds.BST{},
					},
				},
			},
			true,
		},
	}

	for _, c := range cases {
		if !c.expected {
			t.Errorf("E: Placeholder")
		}
	}
}
