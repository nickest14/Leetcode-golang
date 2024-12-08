// https://leetcode.com/problems/minimum-limit-of-balls-in-a-bag/
// Level: Medium

package leetcode

func minimumSize(nums []int, maxOperations int) int {
	low, high := 1, 0
	for _, num := range nums {
		if num > high {
			high = num
		}
	}

	for low < high {
		mid := (low + high) / 2
		ops := 0
		for _, num := range nums {
			ops += (num - 1) / mid
		}
		if ops <= maxOperations {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return high
}
