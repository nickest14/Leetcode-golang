package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	ratings []int
}

type question135 struct {
	params
	ans int
}

func Test_Problem135(t *testing.T) {
	qs := []question135{
		{
			params{ratings: []int{1, 0, 2}},
			5,
		},
		{
			params{ratings: []int{1, 2, 2}},
			4,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, candy(p.ratings), ans)
	}
}
