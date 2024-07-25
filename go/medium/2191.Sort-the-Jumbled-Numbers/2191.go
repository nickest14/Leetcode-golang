// https://leetcode.com/problems/sort-the-jumbled-numbers/
// Level: Medium

package leetcode

import (
	"sort"
	"strconv"
)

func sortJumbled(mapping []int, nums []int) []int {

	transformToINT := func(num int) int {
		numStr := strconv.Itoa(num)
		mappedStr := ""
		for _, c := range numStr {
			digit, _ := strconv.Atoi(string(c))
			mappedStr += strconv.Itoa(mapping[digit])
		}
		number, _ := strconv.Atoi(mappedStr)
		return number
	}
	numWithIndexes := make(map[int]int)
	numberMap := make(map[int]int)
	for i, num := range nums {
		numWithIndexes[num] = i
		numberMap[num] = transformToINT(num)
	}

	sort.Slice(nums, func(i, j int) bool {
		mappedI := numberMap[nums[i]]
		mappedJ := numberMap[nums[j]]
		if mappedI == mappedJ {
			return numWithIndexes[nums[i]] < numWithIndexes[nums[j]]
		}
		return mappedI < mappedJ
	})

	return nums
}
