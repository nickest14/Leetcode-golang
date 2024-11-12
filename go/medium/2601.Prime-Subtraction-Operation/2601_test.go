package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question2601 struct {
	params
	ans bool
}

func Test_Problem2601(t *testing.T) {
	qs := []question2601{
		{
			params{[]int{4, 9, 6, 10}},
			true,
		},
		{
			params{[]int{6, 8, 11, 12}},
			true,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, primeSubOperation(p.nums), ans)
	}
}
