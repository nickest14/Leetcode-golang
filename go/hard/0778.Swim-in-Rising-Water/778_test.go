package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	grid [][]int
}

type question778 struct {
	params
	ans int
}

func Test_Problem778(t *testing.T) {
	qs := []question778{
		{
			params{[][]int{{0, 2}, {1, 3}}},
			3,
		},
		{
			params{[][]int{{0, 1, 2, 3, 4}, {24, 23, 22, 21, 5}, {12, 13, 14, 15, 16}, {11, 17, 18, 19, 20}, {10, 9, 8, 7, 6}}},
			16,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, swimInWater(p.grid), ans)
	}
}
