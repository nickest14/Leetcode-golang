// https://leetcode.com/problems/number-of-equivalent-domino-pairs/
// Level: Easy

package leetcode

func numEquivDominoPairs(dominoes [][]int) int {
	count := make(map[int]int)
	ans := 0

	for _, domino := range dominoes {
		key := (1 << domino[0]) | (1 << domino[1])
		count[key]++
	}

	for _, c := range count {
		ans += c * (c - 1) / 2
	}

	return ans
}
