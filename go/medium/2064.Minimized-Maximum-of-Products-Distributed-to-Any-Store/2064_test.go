package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n          int
	quantities []int
}

type question2064 struct {
	params
	ans int
}

func Test_Problem2064(t *testing.T) {
	qs := []question2064{
		{
			params{n: 6, quantities: []int{11, 6}},
			3,
		},
		{
			params{n: 7, quantities: []int{15, 10, 10}},
			5,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minimizedMaximum(p.n, p.quantities), ans)
	}
}
