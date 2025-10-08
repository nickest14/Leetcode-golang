// https://leetcode.com/problems/partition-array-such-that-maximum-difference-is-k/
// Level: Medium

package leetcode

import (
	"sort"
)

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	ans := make([]int, len(spells))
	for i, spell := range spells {
		ans[i] = len(potions) - sort.Search(len(potions), func(j int) bool {
			return int64(spell*potions[j]) >= success
		})
	}
	return ans
}
