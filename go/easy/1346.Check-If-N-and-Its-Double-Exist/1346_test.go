package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	arr []int
}

type question1346 struct {
	params
	ans bool
}

func Test_Problem1346(t *testing.T) {
	qs := []question1346{
		{
			params{[]int{10, 2, 5, 3}},
			true,
		},
		{
			params{[]int{3, 1, 7, 11}},
			false,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, checkIfExist(p.arr), ans)
	}
}
