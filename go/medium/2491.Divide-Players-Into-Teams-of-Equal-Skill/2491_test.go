package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	skill []int
}

type question2491 struct {
	params
	ans int64
}

func Test_Problem(t *testing.T) {
	qs := []question2491{
		{
			params{skill: []int{3, 2, 5, 1, 3, 4}},
			22,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, dividePlayers(p.skill), ans)
	}
}
