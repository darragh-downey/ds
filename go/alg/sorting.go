package alg

import "math"

// Selection sort
// Time complexity: O(n^2) as there are two nested loops over the (same) array
// Space complexity: O(1) it never makes more than O(n) swaps and can be useful when memory writes are expensive
func Selection(arr []int) {
	// traverse through all array elements
	for i := 0; i < len(arr); i++ {
		// find the minimum element in remaining unsorted array
		minIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[minIdx] > arr[j] {
				minIdx = j
			}
		}
		// Swap the found minimum element with the first element
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

// StableSelection sort
// Time complexity: O(n^2)
// Space complexity: O(1)
// Fixes unnecessary swaps between elements with the same value
func StableSelection(arr []int) {
	for i := 0; i < len(arr); i++ {
		// find the minimum element in the remaining unsorted array
		minIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[minIdx] > arr[j] {
				minIdx = j
			}
		}
		// Move minimum element at current index
		key := arr[minIdx]
		for minIdx > i {
			arr[minIdx] = arr[minIdx-1]
			minIdx -= 1
		}
		arr[i] = key
	}
}

// Bubble sort
// Time complexity: O(n^2)
// Space complexity: O(1)
func Bubble(arr []int) {
	for i := 0; i < len(arr); i++ {
		// last i elements are already in place
		for j := 0; j < len(arr)-i-1; j++ {
			// traverse the array from 0 to len(arr)-i-1
			// swap if the element found is greater than the next element
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// OptimizedBubble sort
// Worst and Average time complexity: O(n*n) worst case occurs when array is reverse sorted
// Best case time complexity: O(n) occurs when array is already sorted
// Space complexity: O(1)
func OptimizedBubble(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			// traverse the array from 0 to n-i-1.
			// Swap if the element found is greater than the next element
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		// if no two elements were swapped by the inner loop
		if !swapped {
			break
		}
	}
}

// Insertion sort
// Time complexity: O(n^2) worst case if array is in reverse sorted order
// Space complexity: O(1)
func Insertion(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		// Move elements of arr[0..i-1] that are
		// greater than key, to one position ahead
		// of their current position
		j := i - 1
		for j >= 0 && key < arr[j] {
			arr[j+1] = arr[j]
			j -= 1
		}
		arr[j+1] = key
	}
}

func tdmerge(arr []int, l, mid, r int) {
	n1 := mid - l + 1
	n2 := r - mid
	// create temporary arrays
	L, R := make([]int, mid), make([]int, mid)
	// Copy data to temp arrays L[] and R[]
	copy(L, arr[l:n1])
	copy(R, arr[n2:r])
	// initial indexes of L, R and the merged subarray
	i, j, k := 0, 0, l

	for i < n1 && j < n2 {
		if L[i] <= R[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
			j++
		}
		k++
	}
	// Copy the remaining elements of L[], if there are any
	for i < n1 {
		arr[k] = L[i]
		i++
		k++
	}
	// Copy the remaining elements of R[], if there are any
	for j < n2 {
		arr[k] = R[j]
		j++
		k++
	}
}

func TopDownMerge(arr []int, l, r int) {
	if l >= r {
		return
	}
	mid := l + (r-l)/2
	TopDownMerge(arr, l, mid)
	TopDownMerge(arr, mid+1, r)
	tdmerge(arr, l, mid, r)
}

// https://algs4.cs.princeton.edu/22mergesort/MergeBU.java.html
// stably merge arr[lo..mid] with arr[mid+1..hi] using temp[lo..hi]
func bumerge(arr, temp []int, lo, mid, hi int) {
	for k := lo; k <= hi; k++ {
		temp[k] = arr[k]
	}
	i, j := lo, mid+1

	for k := lo; k <= hi; k++ {
		if i > mid {
			arr[k] = temp[j]
			j++
		} else if j > hi {
			arr[k] = temp[i]
			i++
		} else if temp[j] < temp[i] {
			arr[k] = temp[j]
			j++
		} else {
			arr[k] = temp[i]
			i++
		}
	}
}

func BottomUpMerge(arr []int) {
	n := len(arr)
	temp := make([]int, n)

	for length := 1; length < n; length *= 2 {
		for lo := 0; lo < n-length; lo += length + length {
			mid := lo + length - 1
			hi := int(math.Min(float64(lo+length+length-1), float64(n-1)))
			bumerge(arr, temp, lo, mid, hi)
		}
	}
}

// to heapify a subtree rooted with node i which is
// an index in arr []int. n is size of heap
func heapify(arr []int, n, i int) {
	largest := i // initialize largest as root
	l := 2*i + 1 // left = 2*i + 1
	r := 2*i + 2 // right = 2*i + 2

	// if left child is larger than root
	if l < n && arr[l] > arr[largest] {
		largest = l
	}
	// if right child is larger than largest so far
	if r < n && arr[r] > arr[largest] {
		largest = r
	}
	// if largest is not root
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		// recursively heapify the affected subtree
		heapify(arr, n, largest)
	}
}

// Heap sort
func Heap(arr []int, n int) {
	// build heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}
	// One by one extract an element from heap
	for i := n - 1; i > 0; i-- {
		// move current root to end
		arr[0], arr[i] = arr[i], arr[0]

		// call max heapify on the reduced heap
		heapify(arr, i, 0)
	}
}
