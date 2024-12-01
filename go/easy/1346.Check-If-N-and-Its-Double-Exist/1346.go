// https://leetcode.com/problems/check-if-n-and-its-double-exist/
// Level: Easy

package leetcode

import "sort"

func checkIfExist(arr []int) bool {
	sort.Ints(arr)
	binarySearch := func(target int) int {
		low, high := 0, len(arr)-1
		for low <= high {
			mid := (low + high) / 2
			if arr[mid] == target {
				return mid
			} else if arr[mid] < target {
				low = mid + 1
			} else if arr[mid] > target {
				high = mid - 1
			}
		}
		return -1
	}

	for i, val := range arr {
		index := binarySearch(val * 2)
		if index != -1 && index != i {
			return true
		}
	}
	return false
}
