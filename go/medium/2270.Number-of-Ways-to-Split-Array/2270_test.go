package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question2270 struct {
	params
	ans int
}

func Test_Problem2270(t *testing.T) {
	qs := []question2270{
		{
			params{nums: []int{10, 4, -8, 7}},
			2,
		},
		{
			params{nums: []int{2, 3, 1, 0}},
			2,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, waysToSplitArray(p.nums), ans)
	}
}
