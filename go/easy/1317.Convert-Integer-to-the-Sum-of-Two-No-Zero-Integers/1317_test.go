package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n int
}

type question1317 struct {
	params
	ans []int
}

func Test_Problem1317(t *testing.T) {
	qs := []question1317{
		{
			params{2},
			[]int{1, 1},
		},
		{
			params{11},
			[]int{2, 9},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, getNoZeroIntegers(p.n), ans)
	}
}
