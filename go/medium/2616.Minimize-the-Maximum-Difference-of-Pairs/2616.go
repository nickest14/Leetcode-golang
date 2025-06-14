// https://leetcode.com/problems/minimize-the-maximum-difference-of-pairs/
// Level: Medium

package leetcode

import "sort"

func minimizeMax(nums []int, p int) int {
	n := len(nums)
	sort.Ints(nums)
	low, high := 0, nums[n-1]-nums[0]

	for low < high {
		mid := (low + high) / 2
		count := 0
		i := 1

		for i < n && count < n {
			if nums[i]-nums[i-1] <= mid {
				count++
				i += 2
			} else {
				i++
			}
		}

		if count >= p {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}
