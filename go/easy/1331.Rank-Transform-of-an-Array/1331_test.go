package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	arr []int
}

type question1331 struct {
	params
	ans []int
}

func Test_Problem1331(t *testing.T) {
	qs := []question1331{
		{
			params{[]int{40, 10, 20, 30}},
			[]int{4, 1, 2, 3},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, arrayRankTransform(p.arr), ans)
	}
}
