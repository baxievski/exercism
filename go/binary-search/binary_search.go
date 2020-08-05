// Package binarysearch implements a solution for the exercism binary-search challenge
package binarysearch

// SearchInts returns the index in the `SearchInts` whose value is the same as `value`
func SearchInts(sortedSlice []int, value int) int {
	if len(sortedSlice) == 1 {
		if sortedSlice[0] == value {
			return 0
		}
	}
	for start, end := 0, len(sortedSlice)-1; start < end; {
		if sortedSlice[start] == value {
			return start
		}
		if sortedSlice[end] == value {
			return end
		}
		if sortedSlice[start] > value {
			return -1
		}
		if sortedSlice[end] < value {
			return -1
		}
		middle := (start + end) / 2
		if sortedSlice[middle] > value {
			end = middle - 1
			continue
		}
		if sortedSlice[middle] < value {
			start = middle + 1
			continue
		}
		return middle
	}
	return -1
}
