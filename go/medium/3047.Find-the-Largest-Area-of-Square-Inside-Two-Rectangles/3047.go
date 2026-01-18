// https://leetcode.com/problems/find-the-largest-area-of-square-inside-two-rectangles/
// Level: Medium

package leetcode

func largestSquareArea(bottomLeft [][]int, topRight [][]int) int64 {
	n := len(bottomLeft)
	maxSideLength := 0

	for i := 0; i < n-1; i++ {
		left1, bottom1 := bottomLeft[i][0], bottomLeft[i][1]
		right1, top1 := topRight[i][0], topRight[i][1]

		for j := i + 1; j < n; j++ {
			left2, bottom2 := bottomLeft[j][0], bottomLeft[j][1]
			right2, top2 := topRight[j][0], topRight[j][1]

			width := min(right1, right2) - max(left1, left2)
			height := min(top1, top2) - max(bottom1, bottom2)

			if width > 0 && height > 0 {
				sideLength := min(width, height)
				if sideLength > maxSideLength {
					maxSideLength = sideLength
				}
			}
		}
	}

	return int64(maxSideLength * maxSideLength)
}
