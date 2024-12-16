package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	classes       [][]int
	extraStudents int
}

type question1792 struct {
	params
	ans float64
}

func Test_Problem1792(t *testing.T) {
	qs := []question1792{
		// {
		// 	params{classes: [][]int{{1, 2}, {3, 5}, {2, 2}}, extraStudents: 2},
		// 	0.78333,
		// },
		{
			params{classes: [][]int{{2, 4}, {3, 9}, {4, 5}, {2, 10}}, extraStudents: 4},
			0.53485,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maxAverageRatio(p.classes, p.extraStudents), ans)
	}
}
