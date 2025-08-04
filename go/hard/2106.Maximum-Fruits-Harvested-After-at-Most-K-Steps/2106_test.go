package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	fruits   [][]int
	startPos int
	k        int
}

type question2106 struct {
	params
	ans int
}

func Test_Problem2106(t *testing.T) {
	qs := []question2106{
		{
			params{fruits: [][]int{{2, 8}, {6, 3}, {8, 6}}, startPos: 5, k: 4},
			9,
		},
		{
			params{fruits: [][]int{{0, 9}, {4, 1}, {5, 7}, {6, 2}, {7, 4}, {10, 9}}, startPos: 5, k: 4},
			14,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maxTotalFruits(p.fruits, p.startPos, p.k), ans)
	}
}
