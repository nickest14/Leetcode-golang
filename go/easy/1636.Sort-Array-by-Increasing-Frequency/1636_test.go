package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question1636 struct {
	params
	ans []int
}

func Test_Problem1636(t *testing.T) {
	qs := []question1636{
		{
			params{[]int{2, 3, 1, 3, 2}},
			[]int{1, 3, 3, 2, 2},
		},
	}

	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, frequencySort(p.nums), ans)
	}
}
