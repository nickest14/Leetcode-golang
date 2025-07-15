// https://leetcode.com/problems/maximum-matching-of-players-with-trainers/
// Level: Medium

package leetcode

import "sort"

func matchPlayersAndTrainers(players []int, trainers []int) int {
	sort.Ints(players)
	sort.Ints(trainers)

	ans := 0
	for i, j := 0, 0; i < len(players) && j < len(trainers); {
		if players[i] <= trainers[j] {
			ans++
			i++
			j++
		} else {
			j++
		}
	}
	return ans
}
