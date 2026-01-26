package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	arr []int
}

type question1200 struct {
	params
	ans [][]int
}

func Test_Problem1200(t *testing.T) {
	qs := []question1200{
		{
			params{[]int{4, 2, 1, 3}},
			[][]int{{1, 2}, {2, 3}, {3, 4}},
		},
		{
			params{[]int{1, 3, 6, 10, 15}},
			[][]int{{1, 3}},
		},
		{
			params{[]int{3, 8, -10, 23, 19, -4, -14, 27}},
			[][]int{{-14, -10}, {19, 23}, {23, 27}},
		},
	}

	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minimumAbsDifference(p.arr), ans)
	}
}
