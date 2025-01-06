// https://leetcode.com/problems/minimum-number-of-operations-to-move-all-balls-to-each-box/
// Level: Medium

package leetcode

func minOperations(boxes string) []int {
	n := len(boxes)
	ans := make([]int, n)
	prefixCount := 0
	prefixSum := 0

	for i := 0; i < n; i++ {
		ans[i] += prefixCount*i - prefixSum
		if boxes[i] == '1' {
			prefixCount++
			prefixSum += i
		}
	}

	suffixCount := 0
	suffixSum := 0
	for i := n - 1; i >= 0; i-- {
		ans[i] += suffixSum - suffixCount*i
		if boxes[i] == '1' {
			suffixCount++
			suffixSum += i
		}
	}

	return ans
}
