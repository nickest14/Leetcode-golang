package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	height []int
}

type question11 struct {
	params
	ans int
}

func Test_Problem11(t *testing.T) {

	qs := []question11{
		{
			params{height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}},
			49,
		},
		{
			params{height: []int{1, 1}},
			1,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v, [answer]: %v\n", p, maxArea(p.height), ans)
	}
}
