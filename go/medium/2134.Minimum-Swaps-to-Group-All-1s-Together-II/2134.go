// https://leetcode.com/problems/minimum-swaps-to-group-all-1s-together-ii/
// Level: Medium

package leetcode

func minSwaps(nums []int) int {
	n := len(nums)
	k := 0
	for i := 0; i < n; i++ {
		if nums[i] == 1 {
			k++
		}
	}
	cnt := 0
	for i := 0; i < k; i++ {
		if nums[i] == 1 {
			cnt++
		}
	}
	mx := cnt
	for i := k; i < n+k; i++ {
		cnt += nums[i%n]
		cnt -= nums[(i-k+n)%n]
		mx = max(mx, cnt)
	}

	return k - mx
}
