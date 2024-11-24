// https://leetcode.com/problems/rotating-the-box/
// Level: Medium

package leetcode

func rotateTheBox(box [][]byte) [][]byte {
	rows := len(box)
	cols := len(box[0])
	ans := make([][]byte, cols)
	for i := range ans {
		ans[i] = make([]byte, rows)
		for j := range ans[i] {
			ans[i][j] = '.'
		}
	}

	for r := 0; r < rows; r++ {
		i := cols - 1
		for c := cols - 1; c >= 0; c-- {
			if box[r][c] == '#' {
				ans[i][rows-r-1] = '#'
				i--
			} else if box[r][c] == '*' {
				ans[c][rows-r-1] = '*'
				i = c - 1
			}
		}
	}
	return ans
}
