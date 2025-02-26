package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	arr []int
}

type question1524 struct {
	params
	ans int
}

func Test_Problem1524(t *testing.T) {
	qs := []question1524{
		{
			params{arr: []int{1, 3, 5}},
			4,
		},
		{
			params{arr: []int{2, 4, 6}},
			0,
		},
		{
			params{arr: []int{1, 2, 3, 4, 5, 6, 7}},
			16,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, numOfSubarrays(p.arr), ans)
	}
}
