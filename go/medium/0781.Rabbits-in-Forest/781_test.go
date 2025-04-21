package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	answers []int
}

type question781 struct {
	params
	ans int
}

func Test_Problem781(t *testing.T) {
	qs := []question781{
		{
			params{answers: []int{1, 1, 2}},
			5,
		},
		{
			params{answers: []int{10, 10, 10}},
			11,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, numRabbits(p.answers), ans)
	}
}
