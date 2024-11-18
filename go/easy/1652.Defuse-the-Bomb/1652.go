// https://leetcode.com/problems/defuse-the-bomb/
// Level: Easy

package leetcode

func decrypt(code []int, k int) []int {
	n := len(code)
	ans := make([]int, n)
	if k == 0 {
		return ans
	}

	code = append(code, code...)
	for i := 0; i < n; i++ {
		if k > 0 {
			for j := 1; j <= k; j++ {
				ans[i] += code[i+j]
			}
		} else {
			for j := k; j < 0; j++ {
				ans[i] += code[n+i+j]
			}
		}
	}

	return ans
}
