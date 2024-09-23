package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n int
}

type question386 struct {
	params
	ans []int
}

func Test_Problem386(t *testing.T) {

	qs := []question386{
		{
			params{13},
			[]int{1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			params{10},
			[]int{1, 10, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v, [answer]: %v\n", p, lexicalOrder(p.n), ans)
	}
}
