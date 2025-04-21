// https://leetcode.com/problems/count-the-hidden-sequences/
// Level: Medium

package leetcode

func numberOfArrays(differences []int, lower int, upper int) int {
	cur, min, max := 0, 0, 0
	for _, diff := range differences {
		cur += diff
		if cur < min {
			min = cur
		}
		if cur > max {
			max = cur
		}
	}
	ans := (upper - lower) - (max - min) + 1
	if ans < 0 {
		return 0
	}
	return ans
}
