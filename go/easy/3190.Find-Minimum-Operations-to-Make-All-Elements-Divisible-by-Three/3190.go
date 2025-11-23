// https://leetcode.com/problems/find-minimum-operations-to-make-all-elements-divisible-by-three/
// Level: Easy

package leetcode

func minimumOperations(nums []int) int {
	ans := 0
	for _, num := range nums {
		if num%3 > 0 {
			ans++
		}
	}
	return ans
}
