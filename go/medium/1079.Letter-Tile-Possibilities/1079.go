// https://leetcode.com/problems/letter-tile-possibilities/description/
// Level: Medium

package leetcode

func numTilePossibilities(tiles string) int {
	charCount := make([]int, 26)
	for _, c := range tiles {
		charCount[c-'A']++
	}

	var buildChar func() int
	buildChar = func() int {
		totalCount := 0
		for i := 0; i < 26; i++ {
			if charCount[i] > 0 {
				totalCount++
				charCount[i]--
				totalCount += buildChar()
				charCount[i]++
			}
		}
		return totalCount
	}

	return buildChar()
}
