package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	tops    []int
	bottoms []int
}

type question1007 struct {
	params
	ans int
}

func Test_Problem1007(t *testing.T) {
	qs := []question1007{
		{
			params{tops: []int{2, 1, 2, 4, 2, 2}, bottoms: []int{5, 2, 6, 2, 3, 2}},
			2,
		},
		{
			params{tops: []int{3, 5, 1, 2, 3}, bottoms: []int{3, 6, 3, 3, 4}},
			-1,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minDominoRotations(p.tops, p.bottoms), ans)
	}
}
