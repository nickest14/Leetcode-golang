package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	target []int
}

type question1526 struct {
	params
	ans int
}

func Test_Problem1526(t *testing.T) {
	qs := []question1526{
		{
			params{target: []int{1, 2, 3, 2, 1}},
			3,
		},
		{
			params{target: []int{3, 1, 1, 2}},
			4,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minNumberOperations(p.target), ans)
	}
}
