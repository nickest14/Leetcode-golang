package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	differences []int
	lower       int
	upper       int
}

type question2145 struct {
	params
	ans int
}

func Test_Problem2145(t *testing.T) {
	qs := []question2145{
		{
			params{differences: []int{1, -3, 4}, lower: 1, upper: 6},
			2,
		},
		{
			params{differences: []int{3, -4, 5, 1, -2}, lower: -4, upper: 5},
			4,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, numberOfArrays(p.differences, p.lower, p.upper), ans)
	}
}
