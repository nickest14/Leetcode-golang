// https://leetcode.com/problems/maximum-frequency-of-an-element-after-performing-operations-i/
// Level: Medium

package leetcode

func maxFrequency(nums []int, k int, numOperations int) int {
	maxNum := nums[0]
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}
	n := maxNum + k + 2
	freq := make([]int, n)
	for _, num := range nums {
		freq[num]++
	}
	pre := make([]int, n)
	pre[0] = freq[0]
	for i := 1; i < n; i++ {
		pre[i] = pre[i-1] + freq[i]
	}
	ans := 0
	for t := 0; t < n; t++ {
		if freq[t] == 0 && numOperations == 0 {
			continue
		}
		l := t - k
		if l < 0 {
			l = 0
		}
		r := t + k
		if r > n-1 {
			r = n - 1
		}
		total := pre[r]
		if l > 0 {
			total -= pre[l-1]
		}
		adj := total - freq[t]
		val := freq[t] + min(numOperations, adj)
		if val > ans {
			ans = val
		}
	}
	return ans
}
