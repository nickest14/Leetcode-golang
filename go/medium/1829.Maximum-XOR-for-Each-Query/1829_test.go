package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums       []int
	maximumBit int
}

type question1829 struct {
	params
	ans []int
}

func Test_Problem1829(t *testing.T) {
	qs := []question1829{
		{
			params{nums: []int{0, 1, 1, 3}, maximumBit: 2},
			[]int{0, 3, 2, 3},
		},
		{
			params{nums: []int{2, 3, 4, 7}, maximumBit: 3},
			[]int{5, 2, 6, 5},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, getMaximumXor(p.nums, p.maximumBit), ans)
	}
}
