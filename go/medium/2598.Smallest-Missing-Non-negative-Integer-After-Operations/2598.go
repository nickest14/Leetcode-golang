// https://leetcode.com/problems/smallest-missing-non-negative-integer-after-operations/
// Level: Medium

package leetcode

func findSmallestInteger(nums []int, value int) int {
	remainderCount := make([]int, value)

	for _, num := range nums {
		rem := ((num % value) + value) % value
		remainderCount[rem]++
	}

	ans := 0
	for remainderCount[ans%value] > 0 {
		remainderCount[ans%value]--
		ans++
	}

	return ans
}
