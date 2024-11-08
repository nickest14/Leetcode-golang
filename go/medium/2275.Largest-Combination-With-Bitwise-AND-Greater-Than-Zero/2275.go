// https://leetcode.com/problems/largest-combination-with-bitwise-and-greater-than-zero/
// Level: Medium

package leetcode

func largestCombination(candidates []int) int {
	ans := 0
	for i := 0; i < 32; i++ {
		count := 0
		for _, c := range candidates {
			if c&(1<<i) != 0 {
				count++
			}
		}
		if count > ans {
			ans = count
		}
	}
	return ans
}
