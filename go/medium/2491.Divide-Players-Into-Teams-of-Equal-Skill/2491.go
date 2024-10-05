// https://leetcode.com/problems/divide-players-into-teams-of-equal-skill/
// Level: Medium

package leetcode

import "sort"

func dividePlayers(skill []int) int64 {
	sort.Ints(skill)

	teamSkill := skill[0] + skill[len(skill)-1]
	var ans int64 = 0
	for i := 0; i < len(skill)/2; i++ {
		if skill[i]+skill[len(skill)-1-i] != teamSkill {
			return -1
		}
		ans += int64(skill[i]) * int64(skill[len(skill)-1-i])
	}
	return ans
}
