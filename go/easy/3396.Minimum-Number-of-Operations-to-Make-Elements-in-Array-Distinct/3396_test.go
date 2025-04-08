package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question3396 struct {
	params
	ans int
}

func Test_Problem3396(t *testing.T) {
	qs := []question3396{
		{
			params{nums: []int{1, 2, 3, 4, 2, 3, 3, 5, 7}},
			2,
		},
		{
			params{nums: []int{4, 5, 6, 4, 4}},
			2,
		},
		{
			params{nums: []int{6, 7, 8, 9}},
			0,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minimumOperations(p.nums), ans)
	}
}
