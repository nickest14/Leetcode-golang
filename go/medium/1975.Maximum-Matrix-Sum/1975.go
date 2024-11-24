// https://leetcode.com/problems/maximum-matrix-sum/

// Level: Medium

package leetcode

import "math"

func maxMatrixSum(matrix [][]int) int64 {
	negCount := 0
	minValue := int64(math.MaxInt64)
	sumAbs := int64(0)
	for _, row := range matrix {
		for _, val := range row {
			if val < 0 {
				negCount++
			}
			absVal := int64(math.Abs(float64(val)))
			sumAbs += absVal
			if absVal < minValue {
				minValue = absVal
			}
		}
	}
	if negCount%2 != 0 {
		sumAbs -= 2 * minValue
	}

	return sumAbs
}
