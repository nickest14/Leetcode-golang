// https://leetcode.com/problems/number-of-subsequences-that-satisfy-the-given-sum-condition/
// Level: Medium

package leetcode

import "sort"

func numSubseq(nums []int, target int) int {
	mod := 1000000007

	pow := func(base, exp int) int {
		result := 1
		for exp > 0 {
			if exp&1 == 1 {
				result = (result * base) % mod
			}
			base = (base * base) % mod
			exp >>= 1
		}
		return result
	}

	sort.Ints(nums)
	ans := 0
	left, right := 0, len(nums)-1

	for left <= right {
		if nums[left]+nums[right] > target {
			right--
		} else {
			ans += pow(2, right-left)
			left++
		}
	}

	return ans % mod
}
