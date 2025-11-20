// https://leetcode.com/problems/keep-multiplying-found-values-by-two/
// Level: Easy

package leetcode

func findFinalValue(nums []int, original int) int {
	bits := 0
	for _, num := range nums {
		if num%original != 0 {
			continue
		}
		n := num / original
		if n&(n-1) == 0 {
			bits |= n
		}
	}
	digit := bits + 1
	return original * (digit & -digit)
}
