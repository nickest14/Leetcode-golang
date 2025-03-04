// https://leetcode.com/problems/partition-array-according-to-given-pivot/description/
// Level: Medium

package leetcode

func pivotArray(nums []int, pivot int) []int {
	var smaller, equal, larger []int

	for _, num := range nums {
		if num < pivot {
			smaller = append(smaller, num)
		} else if num > pivot {
			larger = append(larger, num)
		} else {
			equal = append(equal, num)
		}
	}

	return append(append(smaller, equal...), larger...)
}
