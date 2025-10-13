package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	skill []int
	mana  []int
}

type question3494 struct {
	params
	ans int64
}

func Test_Problem(t *testing.T) {
	qs := []question3494{
		{
			params{skill: []int{1, 5, 2, 4}, mana: []int{5, 1, 4, 2}},
			110,
		},
		{
			params{skill: []int{1, 1, 1}, mana: []int{1, 1, 1}},
			5,
		},
		{
			params{skill: []int{1, 2, 3, 4}, mana: []int{1, 2}},
			21,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minTime(p.skill, p.mana), ans)
	}
}
