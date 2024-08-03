package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question2134 struct {
	params
	ans int
}

func Test_Problem2134(t *testing.T) {
	qs := []question2134{
		{
			params{[]int{0, 1, 0, 1, 1, 0, 0}},
			1,
		},
		{
			params{[]int{0, 1, 1, 1, 0, 0, 1, 1, 0}},
			2,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minSwaps(p.nums), ans)
	}
}
