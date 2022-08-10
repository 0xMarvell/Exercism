package binarysearch

// SearchInts implements a binary search algorithm
func SearchInts(list []int, key int) int {
	var left, right int = 0, len(list) - 1

	for left <= right {
		middle := (left + right) / 2

		if list[middle] == key {
			return middle
		} else if key > list[middle] {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return -1
}
