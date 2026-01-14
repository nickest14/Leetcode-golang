// https://leetcode.com/problems/separate-squares-i/

// Level: Medium

package leetcode

func separateSquares(squares [][]int) float64 {
	areaBelow := func(y float64) float64 {
		area := 0.0
		for _, square := range squares {
			_, y0, sideLen := square[0], float64(square[1]), float64(square[2])
			if y0 < y {
				h := y - y0
				if h > sideLen {
					h = sideLen
				}
				area += h * sideLen
			}
		}
		return area
	}

	totalArea := 0.0
	maxY := 0.0
	for _, square := range squares {
		sideLen := float64(square[2])
		totalArea += sideLen * sideLen
		y := float64(square[1])
		if y+sideLen > maxY {
			maxY = y + sideLen
		}
	}

	targetArea := totalArea / 2.0
	low, high := 0.0, maxY

	for high-low > 1e-5 {
		mid := (low + high) / 2.0
		if areaBelow(mid) < targetArea {
			low = mid
		} else {
			high = mid
		}
	}

	return high
}
