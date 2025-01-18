package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	derived []int
}

type question2683 struct {
	params
	ans bool
}

func Test_Problem2683(t *testing.T) {
	qs := []question2683{
		{
			params{derived: []int{1, 1, 0}},
			true,
		},
		{
			params{derived: []int{1, 1}},
			true,
		},
		{
			params{derived: []int{1, 0}},
			false,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, doesValidArrayExist(p.derived), ans)
	}
}
