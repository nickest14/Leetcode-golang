package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	matrix [][]int
}

type question1380 struct {
	params
	ans []int
}

func Test_Problem1380(t *testing.T) {
	qs := []question1380{
		// {
		// 	params{matrix: [][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}}},
		// 	[]int{15},
		// },
		{
			params{matrix: [][]int{{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12}}},
			[]int{12},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, luckyNumbers(p.matrix), ans)
	}
}
