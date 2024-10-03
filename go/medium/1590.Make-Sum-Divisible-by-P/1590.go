// https://leetcode.com/problems/make-sum-divisible-by-p/
// Level: Medium

package leetcode

func minSubarray(nums []int, p int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	remainder := total % p
	if remainder == 0 {
		return 0
	}

	findSmallestRemainder := func() int {
		precfixSum := 0
		minLength := len(nums)
		prefixMap := map[int]int{0: -1}
		for i, num := range nums {
			precfixSum += num
			targetRemainder := (precfixSum%p - remainder + p) % p
			if index, exists := prefixMap[targetRemainder]; exists {
				minLength = min(minLength, i-index)
			}
			prefixMap[precfixSum%p] = i

		}
		return minLength
	}

	ans := findSmallestRemainder()
	if ans < len(nums) {
		return ans
	} else {
		return -1
	}
}
