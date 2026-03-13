// https://leetcode.com/problems/minimum-number-of-seconds-to-make-mountain-height-zero/
// Level: Medium

package leetcode

import "math"

func minNumberOfSeconds(mountainHeight int, workerTimes []int) int64 {
	canFinish := func(time int) bool {
		total := 0
		for _, w := range workerTimes {
			k := int((math.Sqrt(1+8*float64(time)/float64(w)) - 1) / 2)
			total += k
			if total >= mountainHeight {
				return true
			}
		}
		return false
	}

	maxW := 0
	for _, w := range workerTimes {
		if w > maxW {
			maxW = w
		}
	}
	hi := maxW * mountainHeight * (mountainHeight + 1) / 2
	lo := 0
	for lo < hi {
		mid := (lo + hi) / 2
		if canFinish(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}
