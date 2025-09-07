// https://leetcode.com/problems/find-n-unique-integers-sum-up-to-zero/
// Level: Easy

package leetcode

func sumZero(n int) []int {
	ans := make([]int, n)
	ans[0] = n * (1 - n) / 2
	for i := 1; i < n; i++ {
		ans[i] = i
	}

	return ans
}
