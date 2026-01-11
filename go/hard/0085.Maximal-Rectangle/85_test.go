package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	matrix [][]byte
}

type question85 struct {
	params
	ans int
}

func Test_Problem85(t *testing.T) {
	qs := []question85{
		{
			params{matrix: [][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}},
			6,
		},
		{
			params{matrix: [][]byte{{'0'}}},
			0,
		},
		{
			params{matrix: [][]byte{{'1'}}},
			1,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maximalRectangle(p.matrix), ans)
	}
}
