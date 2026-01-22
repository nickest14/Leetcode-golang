package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question3507 struct {
	params
	ans int
}

func Test_Problem3507(t *testing.T) {
	qs := []question3507{
		{
			params{nums: []int{5, 2, 3, 1}},
			2,
		},
		{
			params{nums: []int{1, 2, 2}},
			0,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minimumPairRemoval(p.nums), ans)
	}
}
