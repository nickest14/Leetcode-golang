// https://leetcode.com/problems/minimum-operations-to-make-array-elements-zero/

// Level: Hard

package leetcode

func minOperations(queries [][]int) int64 {
	ans := 0
	for _, query := range queries {
		left, right := query[0], query[1]
		totalSteps := 0
		layerStart := 1
		layerDepth := 1

		for layerStart <= right {
			start := max(left, layerStart)
			end := min(right, layerStart*4-1)
			if end >= start {
				totalSteps += (end - start + 1) * layerDepth
			}
			layerStart *= 4
			layerDepth++
		}
		ans += (totalSteps / 2) + (totalSteps % 2)
	}
	return int64(ans)
}
