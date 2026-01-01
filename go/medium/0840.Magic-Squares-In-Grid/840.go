// https://leetcode.com/problems/magic-squares-in-grid/
// Level: Medium

package leetcode

func numMagicSquaresInside(grid [][]int) int {
	x, y := len(grid), len(grid[0])

	isMagicSquare := func(r, c int) bool {
		s := make([]int, 9)
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				s[i*3+j] = grid[r+i][c+j]
			}
		}

		seen := make([]bool, 10)
		for _, v := range s {
			if v < 1 || v > 9 || seen[v] {
				return false
			}
			seen[v] = true
		}

		return s[0]+s[1]+s[2] == 15 && s[3]+s[4]+s[5] == 15 && s[6]+s[7]+s[8] == 15 &&
			s[0]+s[3]+s[6] == 15 && s[1]+s[4]+s[7] == 15 && s[2]+s[5]+s[8] == 15 &&
			s[0]+s[4]+s[8] == 15 && s[2]+s[4]+s[6] == 15
	}

	ans := 0
	for i := 0; i < x-2; i++ {
		for j := 0; j < y-2; j++ {
			if isMagicSquare(i, j) {
				ans++
			}
		}
	}
	return ans
}
