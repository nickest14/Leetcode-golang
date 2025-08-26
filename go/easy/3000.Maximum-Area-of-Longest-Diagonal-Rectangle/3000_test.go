package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	dimensions [][]int
}

type question3000 struct {
	params
	ans int
}

func Test_Problem3000(t *testing.T) {
	qs := []question3000{
		{
			params{[][]int{{9, 3}, {8, 6}}},
			48,
		},
		{
			params{[][]int{{3, 4}, {4, 3}}},
			12,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, areaOfMaxDiagonal(p.dimensions), ans)
	}
}
