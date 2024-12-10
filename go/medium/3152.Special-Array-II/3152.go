// https://leetcode.com/problems/special-array-ii/
// Level: Medium

package leetcode

func isArraySpecial(nums []int, queries [][]int) []bool {
	prefix := make([]int, len(nums))
	prefix[0] = 0
	prev := nums[0]%2 == 0 // even: true, odd: false
	for i := 1; i < len(nums); i++ {
		parity := nums[i]%2 == 0
		if prev == parity {
			prefix[i] = prefix[i-1] + 1
		} else {
			prefix[i] = prefix[i-1]
		}
		prev = parity
	}

	ans := []bool{}
	for _, q := range queries {
		ans = append(ans, prefix[q[1]] == prefix[q[0]])
	}
	return ans
}
