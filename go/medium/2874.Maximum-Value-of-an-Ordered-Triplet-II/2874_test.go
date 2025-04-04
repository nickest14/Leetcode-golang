package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question2874 struct {
	params
	ans int64
}

func Test_Problem2874(t *testing.T) {
	qs := []question2874{
		{
			params{nums: []int{12, 6, 1, 2, 7}},
			77,
		},
		{
			params{nums: []int{1, 10, 3, 4, 19}},
			133,
		},
		{
			params{nums: []int{1, 2, 3}},
			0,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maximumTripletValue(p.nums), ans)
	}
}
