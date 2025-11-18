// https://leetcode.com/problems/check-if-all-1s-are-at-least-length-k-places-away/
// Level: Easy

package leetcode

func kLengthApart(nums []int, k int) bool {
	prev := -1
	for i, num := range nums {
		if num == 1 {
			if prev != -1 && i-prev <= k {
				return false
			}
			prev = i
		}
	}
	return true
}
