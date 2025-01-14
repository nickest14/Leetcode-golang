// https://leetcode.com/problems/find-the-prefix-common-array-of-two-arrays/
// Level: Medium

package leetcode

func findThePrefixCommonArray(A []int, B []int) []int {
	ans := make([]int, len(A))
	n := len(A)
	freq := make([]int, n+1)
	common := 0

	for i := 0; i < n; i++ {
		freq[A[i]]++
		if freq[A[i]] == 2 {
			common++
		}
		freq[B[i]]++
		if freq[B[i]] == 2 {
			common++
		}
		ans[i] = common
	}
	return ans
}
