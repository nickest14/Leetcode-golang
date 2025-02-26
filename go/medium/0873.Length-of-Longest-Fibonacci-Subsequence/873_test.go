package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	arr []int
}

type question873 struct {
	params
	ans int
}

func Test_Problem873(t *testing.T) {

	qs := []question873{
		{
			params{[]int{1, 2, 3, 4, 5, 6, 7, 8}},
			5,
		},
		{
			params{[]int{1, 3, 7, 11, 12, 14, 18}},
			3,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v, [answer]: %v\n", p, lenLongestFibSubseq(p.arr), ans)
	}
}
