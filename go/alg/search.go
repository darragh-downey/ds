package alg

import "math"

// Sequential search
// Time complexity: O(n)
// Space complexity: O(1)
func Sequential(arr []int, key int) int {
	for k, i := range arr {
		if i == key {
			return k
		}
	}
	return -1
}

// SequentialImproved search
// Time complexity: O(n)
// Worst case time complexity:
//       1. if element found at last index - O(n) to O(1)
//       2. if element not found - O(n) to O(n/2)
func SequentialImproved(arr []int, key int) int {
	left, right := 0, len(arr)-1
	// run loop from 0 to right
	for left <= right {
		// if key is found with left pointer
		if arr[left] == key {
			return left
		}
		// if key is found with right pointer
		if arr[right] == key {
			return right
		}
		left++
		right--
	}

	// key not found
	return -1
}

// Binary search (iterative)
// Time complexity: O(log n)
// Space complexity: O(1)
func Binary(arr []int, key int) int {
	l, r := 0, len(arr)-1

	for l <= r {
		m := l + (r-l)/2

		if arr[m] == key {
			// found element
			return m
		}

		if arr[m] < key {
			// key is larger, ignore left sub-array
			l = m + 1
		} else {
			// key is smaller, ignore right sub-array
			r = m - 1
		}
	}
	// not found
	return -1
}

// RecursiveBinary search
// Time complexity: O(log n)
// Recursive call stack space: O(log n)
func RecursiveBinary(arr []int, left, right, key int) int {
	if right >= left {
		mid := left + (right-left)/2

		// found element; return
		if key == arr[mid] {
			return mid
		}

		// key is smaller than the element at mid
		// search left sub-array
		if arr[mid] > key {
			return RecursiveBinary(arr, left, mid-1, key)
		}

		// key is greater than the element at mid
		// search the right sub-array
		return RecursiveBinary(arr, mid+1, right, key)
	}
	return -1
}

// Exponential (binary) search
// Use cases:
//     1. Useful for unbounded searches where the size of the array is infinite
//     2. Works better than binary search for bounded arrays, and when the element is closer to the first element
//
// Time complexity: O(log n)
//
// Space complexity: O(log n) for recursive binary search, O(1) for iterative binary search
func Exponential(arr []int, n, key int) int {
	if arr[0] == key {
		return 0
	}

	i := 1
	for i < n && arr[i] <= key {
		i *= 2
	}

	m := 0
	if i < n-1 {
		m = i
	} else {
		m = n - 1
	}
	return RecursiveBinary(arr, i/2, m, key)
}

// Interpol search (Interpolation search)
// Marginally better than Binary search for a uniformly distributed array
// Time complexity: O(log(log(n)))
func Interpol(arr []int, lo, hi, key int) int {
	if lo <= hi && key >= arr[lo] && key <= arr[hi] {
		// The idea of formula is to return higher value of pos
		// when element to be searched is closer to arr[hi].
		// And smaller value when closer to arr[lo]
		pos := lo + int(math.Floor(float64((hi-lo)/(arr[hi]-arr[lo])*(key-arr[lo]))))

		if arr[pos] == key {
			return pos
		}

		if arr[pos] < key {
			return Interpol(arr, pos+1, hi, key)
		}

		if arr[pos] > key {
			return Interpol(arr, lo, pos-1, key)
		}
	}
	return -1
}
