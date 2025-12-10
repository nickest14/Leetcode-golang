// https://leetcode.com/problems/count-special-triplets/
// Level: Medium

package leetcode

func specialTriplets(nums []int) int {
	mod := 1000000007
	maxVal := 0
	for _, x := range nums {
		if x > maxVal {
			maxVal = x
		}
	}
	maxVal *= 2

	freqPrev := make([]int, maxVal+1)
	freqNext := make([]int, maxVal+1)

	for _, x := range nums {
		freqNext[x]++
	}

	ans := 0
	for _, x := range nums {
		freqNext[x]--
		t := x * 2
		ans = (ans + freqPrev[t]*freqNext[t]) % mod
		freqPrev[x]++
	}
	return ans
}
