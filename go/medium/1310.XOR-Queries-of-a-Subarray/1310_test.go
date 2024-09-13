package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	arr     []int
	queries [][]int
}

type question1310 struct {
	params
	ans []int
}

func Test_Problem1310(t *testing.T) {
	qs := []question1310{
		{
			params{arr: []int{1, 3, 4, 8}, queries: [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}}},
			[]int{2, 7, 14, 8},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, xorQueries(p.arr, p.queries), ans)
	}
}
