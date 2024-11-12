// https://leetcode.com/problems/shortest-subarray-with-or-at-least-k-ii/
// Level: Medium

package leetcode

func minimumSubarrayLength(nums []int, k int) int {
	if k == 0 {
		return 1
	}
	ans := len(nums) + 1
	count := make([]int, 32)
	val := 0
	start := 0
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		val |= num
		ibit := 0
		for num > 0 {
			count[ibit] += num & 1
			num >>= 1
			ibit++
		}
		for val >= k && start < len(nums) {
			if i-start+1 < ans {
				ans = i - start + 1
			}
			num := nums[start]
			start++
			ibit = 0
			for num > 0 {
				count[ibit] -= num & 1
				if count[ibit] == 0 {
					val &= ^(1 << ibit)
				}
				num >>= 1
				ibit++
			}
		}
	}
	if ans == len(nums)+1 {
		return -1
	}
	return ans
}
