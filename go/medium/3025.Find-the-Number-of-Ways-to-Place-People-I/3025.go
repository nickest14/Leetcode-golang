// https://leetcode.com/problems/find-the-number-of-ways-to-place-people-i/
// Level: Medium

package leetcode

import "sort"

func numberOfPairs(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] > points[j][1]
		}
		return points[i][0] < points[j][0]
	})

	ans := 0
	n := len(points)

	for i := 0; i < n; i++ {
		upperY := points[i][1]
		lowerYLimit := -1 << 31
		for j := i + 1; j < n; j++ {
			currentY := points[j][1]
			if currentY <= upperY && currentY > lowerYLimit {
				ans++
				lowerYLimit = currentY
				if currentY == upperY {
					break
				}
			}
		}
	}
	return ans
}
