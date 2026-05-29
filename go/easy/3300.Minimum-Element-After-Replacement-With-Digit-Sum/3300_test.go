package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question3300 struct {
	params
	ans int
}

func Test_Problem3300(t *testing.T) {
	qs := []question3300{
		{
			params{nums: []int{10, 12, 13, 14}},
			1,
		},
		{
			params{nums: []int{1, 2, 3, 4}},
			1,
		},
		{
			params{nums: []int{999, 19, 199}},
			10,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minElement(p.nums), ans)
	}
}
