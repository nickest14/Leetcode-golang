package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	values []int
}

type question1039 struct {
	params
	ans int
}

func Test_Problem1039(t *testing.T) {
	qs := []question1039{
		{
			params{values: []int{1, 2, 3}},
			6,
		},
		{
			params{values: []int{3, 7, 4, 5}},
			144,
		},
		{
			params{values: []int{1, 3, 1, 4, 1, 5}},
			13,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minScoreTriangulation(p.values), ans)
	}
}
