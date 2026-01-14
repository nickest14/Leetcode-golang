// https://leetcode.com/problems/minimum-time-visiting-all-points/
// Level: Easy

package leetcode

func minTimeToVisitAllPoints(points [][]int) int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	time := 0
	for i := 1; i < len(points); i++ {
		time += max(abs(points[i][0]-points[i-1][0]), abs(points[i][1]-points[i-1][1]))
	}
	return time
}
