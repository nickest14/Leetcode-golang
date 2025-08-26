// https://leetcode.com/problems/maximum-area-of-longest-diagonal-rectangle/
// Level: Easy

package leetcode

func areaOfMaxDiagonal(dimensions [][]int) int {
	maxArea, maxDiag := 0, 0

	for _, d := range dimensions {
		x, y := d[0], d[1]
		currDiag := x*x + y*y

		if currDiag > maxDiag || (currDiag == maxDiag && x*y > maxArea) {
			maxDiag = currDiag
			maxArea = x * y
		}
	}
	return maxArea
}
