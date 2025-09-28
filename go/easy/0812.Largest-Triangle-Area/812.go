// https://leetcode.com/problems/largest-triangle-area/
// Level: Easy

package leetcode

import "math"

func largestTriangleArea(points [][]int) float64 {
	n := len(points)
	ans := 0.0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				x1, y1 := points[i][0], points[i][1]
				x2, y2 := points[j][0], points[j][1]
				x3, y3 := points[k][0], points[k][1]

				area := 0.5 * math.Abs(float64(x1*(y2-y3)+x2*(y3-y1)+x3*(y1-y2)))
				if area > ans {
					ans = area
				}
			}
		}
	}
	return ans
}
