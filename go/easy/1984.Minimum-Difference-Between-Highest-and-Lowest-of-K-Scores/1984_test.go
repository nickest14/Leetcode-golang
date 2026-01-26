package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
	k    int
}

type question1984 struct {
	params
	ans int
}

func Test_Problem1984(t *testing.T) {
	qs := []question1984{
		{
			params{nums: []int{90}, k: 1},
			0,
		},
		{
			params{nums: []int{9, 4, 1, 7}, k: 2},
			2,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minimumDifference(p.nums, p.k), ans)
	}
}
