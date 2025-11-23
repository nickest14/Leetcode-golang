package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question3190 struct {
	params
	ans int
}

func Test_Problem3190(t *testing.T) {
	qs := []question3190{
		{
			params{nums: []int{1, 2, 3, 4}},
			3,
		},
		{
			params{nums: []int{3, 6, 9}},
			0,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minimumOperations(p.nums), ans)
	}
}
