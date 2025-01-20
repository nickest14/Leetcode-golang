package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	heightMap [][]int
}

type question407 struct {
	params
	ans int
}

func Test_Problem407(t *testing.T) {
	qs := []question407{
		{
			params{[][]int{{1, 4, 3, 1, 3, 2}, {3, 2, 1, 3, 2, 4}, {2, 3, 3, 2, 3, 1}}},
			4,
		},
		{
			params{[][]int{{3, 3, 3, 3, 3}, {3, 2, 2, 2, 3}, {3, 2, 1, 2, 3}, {3, 2, 2, 2, 3}, {3, 3, 3, 3, 3}}},
			10,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, trapRainWater(p.heightMap), ans)
	}
}
