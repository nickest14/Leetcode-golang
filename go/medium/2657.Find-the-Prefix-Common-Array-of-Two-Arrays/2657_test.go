package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	A []int
	B []int
}

type question2657 struct {
	params
	ans []int
}

func Test_Problem2657(t *testing.T) {
	qs := []question2657{
		{
			params{A: []int{1, 3, 2, 4}, B: []int{3, 1, 2, 4}},
			[]int{0, 2, 3, 4},
		},
		{
			params{A: []int{2, 3, 1}, B: []int{3, 1, 2}},
			[]int{0, 1, 3},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, findThePrefixCommonArray(p.A, p.B), ans)
	}
}
