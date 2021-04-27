package util

// See https://github.com/golang/go/wiki/SliceTricks for more tricks

func Insert(arr []int, val, idx int) []int {
	s := append(arr, 0)
	copy(s[idx+1:], s[idx:])
	s[idx] = val
	return s
}

func Delete(arr []int, idx int) {
	copy(arr[idx:], arr[idx+1:])
	arr[len(arr)-1] = 0
	arr = arr[:len(arr)-1]
}

func Cut(arr []int, i, j int) {
	arr = append(arr[:i], arr[j:]...)
}
