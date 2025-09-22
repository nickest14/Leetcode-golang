package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question3005 struct {
	params
	ans int
}

func Test_Problem3005(t *testing.T) {
	qs := []question3005{
		{
			params{nums: []int{1, 2, 2, 3, 1, 4}},
			4,
		},
		{
			params{nums: []int{1, 2, 3, 4, 5}},
			5,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maxFrequencyElements(p.nums), ans)
	}
}
