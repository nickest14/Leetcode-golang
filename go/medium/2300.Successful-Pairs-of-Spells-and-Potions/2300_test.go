package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	spells  []int
	potions []int
	success int64
}

type question2300 struct {
	params
	ans []int
}

func Test_Problem2300(t *testing.T) {
	qs := []question2300{
		{
			params{spells: []int{5, 1, 3}, potions: []int{1, 2, 3, 4, 5}, success: 7},
			[]int{4, 0, 3},
		},
		{
			params{spells: []int{3, 1, 2}, potions: []int{8, 5, 8}, success: 16},
			[]int{2, 0, 2},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, successfulPairs(p.spells, p.potions, p.success), ans)
	}
}
