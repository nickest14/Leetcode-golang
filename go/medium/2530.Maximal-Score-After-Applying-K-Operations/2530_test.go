package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
	k    int
}

type question2530 struct {
	params
	ans int64
}

func Test_Problem2530(t *testing.T) {
	qs := []question2530{
		{
			params{nums: []int{1, 10, 3, 3, 3}, k: 3},
			17,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maxKelements(p.nums, p.k), ans)
	}
}
