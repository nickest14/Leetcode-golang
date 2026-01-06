package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question1390 struct {
	params
	ans int
}

func Test_Problem(t *testing.T) {
	qs := []question1390{
		{
			params{nums: []int{21, 4, 7}},
			32,
		},
		{
			params{nums: []int{21, 21}},
			64,
		},
		{
			params{nums: []int{1, 2, 3, 4, 5}},
			0,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, sumFourDivisors(p.nums), ans)
	}
}
