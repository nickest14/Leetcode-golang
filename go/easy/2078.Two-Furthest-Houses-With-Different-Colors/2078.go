// https://leetcode.com/problems/two-furthest-houses-with-different-colors/
// Level: Easy

package leetcode

func maxDistance(colors []int) int {
	n := len(colors)
	left, right := 0, n-1

	for i := 0; i < n; i++ {
		if colors[i] != colors[n-1] {
			left = i
			break
		}
	}

	for i := n - 1; i >= 0; i-- {
		if colors[i] != colors[0] {
			right = i
			break
		}
	}

	distFromEnd := n - 1 - left
	distFromStart := right
	if distFromEnd > distFromStart {
		return distFromEnd
	}
	return distFromStart
}
