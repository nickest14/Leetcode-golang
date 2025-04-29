package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
	k    int
}

type question2962 struct {
	params
	ans int64
}

func Test_Problem2962(t *testing.T) {
	qs := []question2962{
		// {
		// 	params{nums: []int{1, 3, 2, 3, 3}, k: 2},
		// 	6,
		// },
		// {
		// 	params{nums: []int{1, 4, 2, 1}, k: 3},
		// 	0,
		// },
		{
			params{nums: []int{61, 23, 38, 23, 56, 40, 82, 56, 82, 82, 82, 70, 8, 69, 8, 7, 19, 14, 58, 42, 82, 10, 82, 78, 15, 82}, k: 2},
			224,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, countSubarrays(p.nums, p.k), ans)
	}
}
