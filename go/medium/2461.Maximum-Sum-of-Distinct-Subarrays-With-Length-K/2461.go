// https://leetcode.com/problems/maximum-sum-of-distinct-subarrays-with-length-k/
// Level: Medium

package leetcode

func maximumSubarraySum(nums []int, k int) int64 {
	n := len(nums)
	elements := make(map[int64]bool)
	currentSum := int64(0)
	ans := int64(0)
	begin := 0

	for end := 0; end < n; end++ {
		val := int64(nums[end])
		if !elements[val] {
			currentSum += val
			elements[val] = true

			if end-begin+1 == k {
				if currentSum > ans {
					ans = currentSum
				}
				currentSum -= int64(nums[begin])
				delete(elements, int64(nums[begin]))
				begin++
			}
		} else {
			for begin < end && int64(nums[begin]) != val {
				currentSum -= int64(nums[begin])
				delete(elements, int64(nums[begin]))
				begin++
			}
			begin++
		}

	}
	return ans
}
