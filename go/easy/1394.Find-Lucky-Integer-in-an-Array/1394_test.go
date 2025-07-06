package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	arr []int
}

type question1394 struct {
	params
	ans int
}

func Test_Problem1394(t *testing.T) {
	qs := []question1394{
		{
			params{[]int{2, 2, 3, 4}},
			2,
		},
		{
			params{[]int{1, 2, 2, 3, 3, 3}},
			3,
		},
		{
			params{[]int{2, 2, 2, 3, 3}},
			-1,
		},
	}

	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, findLucky(p.arr), ans)
	}
}
