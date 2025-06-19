package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
	k    int
}

type question2294 struct {
	params
	ans int
}

func Test_Problem2294(t *testing.T) {
	qs := []question2294{
		{
			params{nums: []int{3, 6, 1, 2, 5}, k: 2},
			2,
		},
		{
			params{nums: []int{1, 2, 3}, k: 1},
			2,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, partitionArray(p.nums, p.k), ans)
	}
}
