// https://leetcode.com/problems/find-champion-ii/
// Level: Medium

package leetcode

func findChampion(n int, edges [][]int) int {
	undefeated := make([]bool, n)
	for i := range undefeated {
		undefeated[i] = true
	}

	for _, edge := range edges {
		weaker := edge[1]
		undefeated[weaker] = false
	}

	ans := -1
	count := 0
	for team := 0; team < n; team++ {
		if undefeated[team] {
			ans = team
			count++
			if count > 1 {
				return -1
			}
		}
	}

	return ans
}
