// https://leetcode.com/problems/maximum-fruits-harvested-after-at-most-k-steps/
// Level: Hard

package leetcode

func maxTotalFruits(fruits [][]int, startPos int, k int) int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	minSteps := func(left, right, start int) int {
		a := abs(start-left) + (right - left)
		b := abs(start-right) + (right - left)
		if a < b {
			return a
		}
		return b
	}

	left, sum, maxFruits := 0, 0, 0
	for right := 0; right < len(fruits); right++ {
		sum += fruits[right][1]
		for left <= right && minSteps(fruits[left][0], fruits[right][0], startPos) > k {
			sum -= fruits[left][1]
			left++
		}
		if sum > maxFruits {
			maxFruits = sum
		}
	}

	return maxFruits
}
