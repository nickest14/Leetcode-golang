// https://leetcode.com/problems/number-of-sub-arrays-with-odd-sum/
// Level: Medium

package leetcode

func numOfSubarrays(arr []int) int {
	mod := int(1e9 + 7)
	oddCount, evenCount := 0, 1
	prefixSum, ans := 0, 0

	for _, a := range arr {
		prefixSum += a
		if prefixSum%2 == 1 {
			ans += evenCount
			oddCount++
		} else {
			ans += oddCount
			evenCount++
		}
	}
	return ans % mod
}
