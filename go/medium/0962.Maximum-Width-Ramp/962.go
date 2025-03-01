// https://leetcode.com/problems/maximum-width-ramp/
// Level: Medium

package leetcode

import "sort"

func maxWidthRamp(nums []int) int {
	arr := make([][2]int, len(nums))
	for i, num := range nums {
		arr[i] = [2]int{num, i}
	}

	sort.Slice(arr, func(i, j int) bool {
		if arr[i][0] == arr[j][0] {
			return arr[i][1] < arr[j][1]
		}
		return arr[i][0] < arr[j][0]
	})

	ans := 0
	minIndex := len(nums)
	for _, pair := range arr {
		index := pair[1]
		ans = max(ans, index-minIndex)
		minIndex = min(minIndex, index)
	}
	return ans
}
