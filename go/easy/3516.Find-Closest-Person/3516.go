// https://leetcode.com/problems/find-closest-person/
// Level: Easy

package leetcode

func findClosest(x int, y int, z int) int {
	diffX := abs(x - z)
	diffY := abs(y - z)
	if diffX < diffY {
		return 1
	} else if diffX > diffY {
		return 2
	} else {
		return 0
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
