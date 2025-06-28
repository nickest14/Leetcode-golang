// https://leetcode.com/problems/find-all-k-distant-indices-in-an-array/
// Level: Easy

package leetcode

func findKDistantIndices(nums []int, key int, k int) []int {
	var ans []int
	n := len(nums)
	right := 0

	for i, num := range nums {
		if num == key {
			left := max(right, i-k)
			right = min(n, i+k+1)

			for j := left; j < right; j++ {
				ans = append(ans, j)
			}
		}
	}
	return ans

}
