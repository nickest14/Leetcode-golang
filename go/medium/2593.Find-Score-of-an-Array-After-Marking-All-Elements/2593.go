// https://leetcode.com/problems/find-score-of-an-array-after-marking-all-elements/
// Level: Medium

package leetcode

import "sort"

func findScore(nums []int) int64 {
	ans := 0
	seen := make(map[int]bool)
	n := len(nums)

	type pair struct {
		index int
		value int
	}
	queue := make([]pair, n)
	for i, num := range nums {
		queue[i] = pair{index: i, value: num}
	}
	sort.Slice(queue, func(i, j int) bool {
		if queue[i].value == queue[j].value {
			return queue[i].index < queue[j].index
		}
		return queue[i].value < queue[j].value
	})

	for _, pair := range queue {
		idx, num := pair.index, pair.value
		if seen[idx] {
			continue
		}
		ans += num
		seen[idx] = true
		if idx > 0 {
			seen[idx-1] = true
		}
		if idx < n-1 {
			seen[idx+1] = true
		}
	}

	return int64(ans)
}
