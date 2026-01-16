// https://leetcode.com/problems/maximum-square-area-by-removing-fences-from-a-field/
// Level: Medium

package leetcode

import "sort"

func maximizeSquareArea(m int, n int, hFences []int, vFences []int) int {
	mod := 1000000007

	hFences = append(hFences, 1, m)
	vFences = append(vFences, 1, n)

	sort.Ints(hFences)
	sort.Ints(vFences)

	hGaps := make(map[int]bool)
	for i := 0; i < len(hFences); i++ {
		for j := i + 1; j < len(hFences); j++ {
			hGaps[hFences[j]-hFences[i]] = true
		}
	}

	maxCommon := -1
	for i := 0; i < len(vFences); i++ {
		for j := i + 1; j < len(vFences); j++ {
			gap := vFences[j] - vFences[i]
			if hGaps[gap] {
				if gap > maxCommon {
					maxCommon = gap
				}
			}
		}
	}

	if maxCommon == -1 {
		return -1
	}
	return (maxCommon * maxCommon) % mod
}
