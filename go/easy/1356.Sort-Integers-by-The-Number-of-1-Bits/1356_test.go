package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	arr []int
}

type question1356 struct {
	params
	ans []int
}

func Test_Problem1356(t *testing.T) {
	qs := []question1356{
		{
			params{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8}},
			[]int{0, 1, 2, 4, 8, 3, 5, 6, 7},
		},
		{
			params{[]int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}},
			[]int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024},
		},
	}

	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, sortByBits(p.arr), ans)
	}
}
