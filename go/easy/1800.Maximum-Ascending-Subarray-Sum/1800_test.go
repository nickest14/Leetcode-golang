package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question3105 struct {
	params
	ans int
}

func Test_Problem3105(t *testing.T) {
	qs := []question3105{
		{
			params{[]int{10, 20, 30, 5, 10, 50}},
			65,
		},
		{
			params{[]int{10, 20, 30, 40, 50}},
			150,
		},
		{
			params{[]int{12, 17, 15, 13, 10, 11, 12}},
			33,
		},
	}

	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maxAscendingSum(p.nums), ans)
	}
}
