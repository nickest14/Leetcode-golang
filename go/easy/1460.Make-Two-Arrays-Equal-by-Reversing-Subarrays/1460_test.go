package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	target []int
	arr    []int
}

type question2134 struct {
	params
	ans bool
}

func Test_Problem2134(t *testing.T) {
	qs := []question2134{
		{
			params{[]int{1, 2, 3, 4}, []int{2, 4, 1, 3}},
			true,
		},
	}

	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, canBeEqual(p.target, p.arr), ans)
	}
}
