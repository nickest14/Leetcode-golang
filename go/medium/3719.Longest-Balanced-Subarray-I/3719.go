// https://leetcode.com/problems/longest-balanced-subarray-i/
// Level: Medium

package leetcode

func longestBalanced(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	maxVal := nums[0]
	for _, v := range nums {
		if v > maxVal {
			maxVal = v
		}
	}
	seen := make([]int, maxVal+1)

	ans := 0
	for i := 0; i < n; i++ {
		if n-i <= ans {
			break
		}
		distinctEvenOdd := [2]int{0, 0}
		for j := i; j < n; j++ {
			val := nums[j]
			if seen[val] != i+1 {
				seen[val] = i + 1
				distinctEvenOdd[val&1]++
			}
			if distinctEvenOdd[0] == distinctEvenOdd[1] {
				ans = max(ans, j-i+1)
			}
		}
	}
	return ans
}
