// https://leetcode.com/problems/count-collisions-on-a-road/
// Level: Medium

package leetcode

func countCollisions(directions string) int {
	length := len(directions)
	i, j := 0, length-1
	for i < length && directions[i] == 'L' {
		i++
	}
	for j >= 0 && directions[j] == 'R' {
		j--
	}

	ans := 0
	for k := i; k <= j; k++ {
		if directions[k] != 'S' {
			ans++
		}
	}
	return ans
}
