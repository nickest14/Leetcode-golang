// https://leetcode.com/problems/find-missing-and-repeated-values/
// Level: Easy

package leetcode

func findMissingAndRepeatedValues(grid [][]int) []int {
	n := len(grid)
	numsSet := make(map[int]bool)

	repeated := -1
	for _, row := range grid {
		for _, num := range row {
			if numsSet[num] {
				repeated = num
			}
			numsSet[num] = true
		}
	}

	missing := -1
	for num := 1; num <= n*n; num++ {
		if !numsSet[num] {
			missing = num
			break
		}
	}

	return []int{repeated, missing}
}
