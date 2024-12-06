package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	banned []int
	n      int
	maxSum int
}

type question2554 struct {
	params
	ans int
}

func Test_Problem2554(t *testing.T) {
	qs := []question2554{
		{
			params{banned: []int{1, 6, 5}, n: 5, maxSum: 6},
			2,
		},
		{
			params{banned: []int{1, 2, 3, 4, 5, 6, 7}, n: 8, maxSum: 1},
			0,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maxCount(p.banned, p.n, p.maxSum), ans)
	}
}
