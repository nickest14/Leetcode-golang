// https://leetcode.com/problems/minimum-operations-to-make-array-sum-divisible-by-k/
// Level: Easy

package leetcode

func minOperations(nums []int, k int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	return sum % k
}
