// https://leetcode.com/problems/minimum-operations-to-make-array-values-equal-to-k/
// Level: Easy

package leetcode

func minOperations(nums []int, k int) int {
	visited := make(map[int]bool)
	for _, num := range nums {
		if num < k {
			return -1
		} else if num > k {
			visited[num] = true
		}
	}
	return len(visited)
}
