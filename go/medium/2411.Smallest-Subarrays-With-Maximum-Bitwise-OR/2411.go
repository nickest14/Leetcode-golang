// https://leetcode.com/problems/smallest-subarrays-with-maximum-bitwise-or/
// Level: Medium

package leetcode

func smallestSubarrays(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)

	latest := make([]int, 32)
	for i := range latest {
		latest[i] = -1
	}

	for i := n - 1; i >= 0; i-- {
		farthest := i
		for b := 0; b < 32; b++ {
			if (nums[i]>>b)&1 == 1 {
				latest[b] = i
			}
			if latest[b] != -1 {
				farthest = max(farthest, latest[b])
			}
		}

		ans[i] = farthest - i + 1
	}

	return ans
}
