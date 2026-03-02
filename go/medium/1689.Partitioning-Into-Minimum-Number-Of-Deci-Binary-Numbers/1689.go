// https://leetcode.com/problems/partitioning-into-minimum-number-of-deci-binary-numbers/
// Level: Medium

package leetcode

func minPartitions(n string) int {
	ans := 0
	for _, c := range n {
		if t := int(c - '0'); ans < t {
			ans = t
		}
	}
	return ans
}
