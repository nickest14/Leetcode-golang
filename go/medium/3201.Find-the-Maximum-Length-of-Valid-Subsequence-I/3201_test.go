package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question3201 struct {
	params
	ans int
}

func Test_Problem3201(t *testing.T) {
	qs := []question3201{
		{
			params{nums: []int{1, 2, 3, 4}},
			4,
		},
		{
			params{nums: []int{1, 2, 1, 1, 2, 1, 2}},
			6,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maximumLength(p.nums), ans)
	}
}
