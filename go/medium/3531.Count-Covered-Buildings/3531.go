// https://leetcode.com/problems/count-covered-buildings/
// Level: Medium

package leetcode

func countCoveredBuildings(n int, buildings [][]int) int {
	xMax := make([]int, n+1)
	xMin := make([]int, n+1)
	yMax := make([]int, n+1)
	yMin := make([]int, n+1)

	for i := range xMin {
		xMin[i] = n + 1
	}
	for i := range yMin {
		yMin[i] = n + 1
	}

	for _, building := range buildings {
		x, y := building[0], building[1]
		if x > xMax[y] {
			xMax[y] = x
		}
		if x < xMin[y] {
			xMin[y] = x
		}
		if y > yMax[x] {
			yMax[x] = y
		}
		if y < yMin[x] {
			yMin[x] = y
		}
	}

	ans := 0
	for _, building := range buildings {
		x, y := building[0], building[1]
		if xMin[y] < x && x < xMax[y] && yMin[x] < y && y < yMax[x] {
			ans++
		}
	}
	return ans
}
