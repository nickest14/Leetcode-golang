// https://leetcode.com/problems/lexicographical-numbers/
// Level: Medium

package leetcode

func lexicalOrder(n int) []int {
	ans := []int{}
	cur := 1
	for i := 0; i < n; i++ {
		ans = append(ans, cur)
		if cur*10 <= n {
			cur *= 10
		} else {
			if cur >= n {
				cur /= 10
			}
			cur++
			for cur%10 == 0 {
				cur /= 10
			}
		}
	}
	return ans
}
