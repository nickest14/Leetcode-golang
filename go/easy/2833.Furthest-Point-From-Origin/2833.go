// https://leetcode.com/problems/furthest-point-from-origin/
// Level: Easy

package leetcode

func furthestDistanceFromOrigin(moves string) int {
	left := 0
	right := 0
	distance := 0

	for _, move := range moves {
		if move == 'L' {
			left++
		} else if move == 'R' {
			right++
		} else {
			distance++
		}
	}

	if left > right {
		return left - right + distance
	}
	return right - left + distance
}
