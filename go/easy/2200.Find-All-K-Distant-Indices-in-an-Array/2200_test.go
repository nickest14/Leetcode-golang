package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
	key  int
	k    int
}

type question2200 struct {
	params
	ans []int
}

func Test_Problem2200(t *testing.T) {
	qs := []question2200{
		{
			params{nums: []int{3, 4, 9, 1, 3, 9, 5}, key: 9, k: 1},
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			params{nums: []int{2, 2, 2, 2, 2}, key: 2, k: 2},
			[]int{0, 1, 2, 3, 4},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, findKDistantIndices(p.nums, p.key, p.k), ans)
	}
}
