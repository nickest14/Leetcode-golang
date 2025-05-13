package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	digits []int
}

type question2094 struct {
	params
	ans []int
}

func Test_Problem2094(t *testing.T) {
	qs := []question2094{
		{
			params{[]int{2, 1, 3, 0}},
			[]int{102, 120, 130, 132, 210, 230, 302, 310, 312, 320},
		},
		{
			params{[]int{2, 2, 8, 8, 2}},
			[]int{222, 228, 282, 288, 822, 828, 882},
		},
	}

	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, findEvenNumbers(p.digits), ans)
	}
}
