package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	isWater [][]int
}

type question1765 struct {
	params
	ans [][]int
}

func Test_Problem1765(t *testing.T) {
	qs := []question1765{
		{
			params{[][]int{{0, 1}, {0, 0}}},
			[][]int{{1, 0}, {2, 1}},
		},
		{
			params{[][]int{{0, 0, 1}, {1, 0, 0}, {0, 0, 0}}},
			[][]int{{1, 1, 0}, {0, 1, 1}, {1, 2, 2}},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, highestPeak(p.isWater), ans)
	}
}
