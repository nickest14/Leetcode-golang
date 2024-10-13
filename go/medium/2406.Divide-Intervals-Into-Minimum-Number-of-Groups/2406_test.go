package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	intervals [][]int
}

type question2406 struct {
	params
	ans int64
}

func Test_Problem2406(t *testing.T) {
	qs := []question2406{
		{
			params{[][]int{{5, 10}, {6, 8}, {1, 5}, {2, 3}, {1, 10}}},
			3,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minGroups(p.intervals), ans)
	}
}
