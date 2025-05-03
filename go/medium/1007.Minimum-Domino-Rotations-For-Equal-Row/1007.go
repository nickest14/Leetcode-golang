// https://leetcode.com/problems/minimum-domino-rotations-for-equal-row/
// Level: Medium

package leetcode

func minDominoRotations(tops []int, bottoms []int) int {
	check := func(target int) int {
		topFlips := 0
		bottomFlips := 0
		for i := 0; i < len(tops); i++ {
			t, b := tops[i], bottoms[i]
			if t != target && b != target {
				return -1
			}
			if t != target {
				topFlips++
			} else if b != target {
				bottomFlips++
			}
		}
		return min(topFlips, bottomFlips)
	}

	ans := check(tops[0])
	if ans != -1 || tops[0] == bottoms[0] {
		return ans
	}
	return check(bottoms[0])
}
