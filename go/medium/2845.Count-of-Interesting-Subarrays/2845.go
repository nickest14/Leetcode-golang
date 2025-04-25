// https://leetcode.com/problems/count-of-interesting-subarrays/
// Level: Medium

package leetcode

func countInterestingSubarrays(nums []int, modulo int, k int) int64 {
	ans := 0
	prefixMod := 0
	modPrefix := make(map[int]int)
	modPrefix[0] = 1
	for _, num := range nums {
		if num%modulo == k {
			prefixMod = (prefixMod + 1) % modulo
		}
		needed := (prefixMod - k) % modulo
		if needed < 0 {
			needed += modulo
		}
		ans += modPrefix[needed]
		modPrefix[prefixMod]++
	}

	return int64(ans)
}
