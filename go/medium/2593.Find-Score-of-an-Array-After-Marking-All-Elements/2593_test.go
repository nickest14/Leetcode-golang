package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question2593 struct {
	params
	ans int64
}

func Test_Problem2593(t *testing.T) {

	qs := []question2593{
		{
			params{[]int{2, 1, 3, 4, 5, 2}},
			7,
		},
		{
			params{[]int{2, 3, 5, 1, 3, 2}},
			5,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v, [answer]: %v\n", p, findScore(p.nums), ans)
	}
}
