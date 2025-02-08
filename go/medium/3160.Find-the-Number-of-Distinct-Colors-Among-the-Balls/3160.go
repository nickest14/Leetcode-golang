// https://leetcode.com/problems/find-the-number-of-distinct-colors-among-the-balls/
// Level: Medium

package leetcode

func queryResults(limit int, queries [][]int) []int {
	ans := []int{}
	ballMap := make(map[int]int)
	colorFreq := make(map[int]int)

	for _, query := range queries {
		ball, color := query[0], query[1]
		if oldColor, exists := ballMap[ball]; exists {
			if colorFreq[oldColor] > 1 {
				colorFreq[oldColor]--
			} else {
				delete(colorFreq, oldColor)
			}
		}

		ballMap[ball] = color
		colorFreq[color]++
		ans = append(ans, len(colorFreq))
	}

	return ans
}
