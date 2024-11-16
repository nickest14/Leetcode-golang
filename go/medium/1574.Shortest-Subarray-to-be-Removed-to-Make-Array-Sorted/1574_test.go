package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	arr []int
}

type question1574 struct {
	params
	ans int
}

func Test_Problem1574(t *testing.T) {
	qs := []question1574{
		{
			params{[]int{1, 2, 3, 10, 4, 2, 3, 5}},
			3,
		},
		{
			params{[]int{5, 4, 3, 2, 1}},
			4,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, findLengthOfShortestSubarray(p.arr), ans)
	}
}
