// https://leetcode.com/problems/find-the-maximum-sum-of-node-values/
// Level: Hard

package leetcode

func maximumValueSum(nums []int, k int, edges [][]int) int64 {
	base := 0
	for _, num := range nums {
		base += num
	}

	sumPos := 0
	cntPos := 0
	minPos := 1<<31 - 1
	bestNonPos := -1 << 31

	for _, x := range nums {
		d := (x ^ k) - x
		if d > 0 {
			cntPos++
			sumPos += d
			if d < minPos {
				minPos = d
			}
		} else {
			if d > bestNonPos {
				bestNonPos = d
			}
		}
	}

	if cntPos%2 == 0 {
		return int64(base + sumPos)
	}

	loss := minPos
	if -bestNonPos < loss {
		loss = -bestNonPos
	}

	return int64(base + sumPos - loss)
}
