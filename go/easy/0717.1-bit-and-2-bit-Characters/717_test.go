package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	bits []int
}

type question717 struct {
	params
	ans bool
}

func Test_Problem717(t *testing.T) {
	qs := []question717{
		{
			params{[]int{1, 0, 0}},
			true,
		},
		{
			params{[]int{1, 1, 1, 0}},
			false,
		},
	}

	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, isOneBitCharacter(p.bits), ans)
	}
}
