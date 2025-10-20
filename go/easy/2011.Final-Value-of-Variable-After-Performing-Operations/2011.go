// https://leetcode.com/problems/final-value-of-variable-after-performing-operations/
// Level: Easy

package leetcode

func finalValueAfterOperations(operations []string) int {
	ans := 0
	for _, op := range operations {
		if op[0] == '+' || op[2] == '+' {
			ans++
		} else {
			ans--
		}
	}
	return ans
}
