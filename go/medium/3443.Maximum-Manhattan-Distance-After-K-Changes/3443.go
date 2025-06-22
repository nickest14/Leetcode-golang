// https://leetcode.com/problems/maximum-manhattan-distance-after-k-changes/
// Level: Medium

package leetcode

func maxDistance(s string, k int) int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	ans := 0
	northCount := 0
	southCount := 0
	eastCount := 0
	westCount := 0

	for i, direction := range s {
		switch direction {
		case 'N':
			northCount++
		case 'S':
			southCount++
		case 'E':
			eastCount++
		case 'W':
			westCount++
		}

		currentDistance := abs(northCount-southCount) + abs(eastCount-westCount)
		totalDistance := currentDistance + min(k*2, i+1-currentDistance)
		if totalDistance > ans {
			ans = totalDistance
		}
	}

	return ans
}
