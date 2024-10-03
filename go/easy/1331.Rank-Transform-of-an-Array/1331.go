// https://leetcode.com/problems/rank-transform-of-an-array/
// Level: Easy

package leetcode

import (
	"sort"
)

func arrayRankTransform(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	sortedArr := append([]int{}, arr...)
	sort.Ints(sortedArr)

	rankMap := make(map[int]int)
	rank := 1
	for _, val := range sortedArr {
		if _, exists := rankMap[val]; !exists {
			rankMap[val] = rank
			rank++
		}
	}

	for i, val := range arr {
		arr[i] = rankMap[val]
	}

	return arr
}
