package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	arr []int
	k   int
}

type question1497 struct {
	params
	ans bool
}

func Test_Problem(t *testing.T) {
	qs := []question1497{
		{
			params{arr: []int{1, 2, 3, 4, 5, 10, 6, 7, 8, 9}, k: 5},
			true,
		},
		{
			params{arr: []int{-1, 1, -2, 2, -3, 3, -4, 4}, k: 3},
			true,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, canArrange(p.arr, p.k), ans)
	}
}
