// https://leetcode.com/problems/maximum-subarray-sum-with-length-divisible-by-k/
// Level: Medium

package leetcode

func maxSubarraySum(nums []int, k int) int64 {
	const INF int64 = 1<<62 - 1

	minPrefix := make([]int64, k)
	for i := 0; i < k; i++ {
		minPrefix[i] = INF
	}
	minPrefix[0] = 0

	var prefix int64 = 0
	var ans int64 = -INF

	for i, x := range nums {
		prefix += int64(x)
		mod := (i + 1) % k

		if minPrefix[mod] != INF {
			val := prefix - minPrefix[mod]
			if val > ans {
				ans = val
			}
		}

		if prefix < minPrefix[mod] {
			minPrefix[mod] = prefix
		}
	}

	return ans
}
