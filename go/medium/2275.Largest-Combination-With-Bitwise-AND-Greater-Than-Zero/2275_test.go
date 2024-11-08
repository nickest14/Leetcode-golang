package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	candidates []int
}

type question2275 struct {
	params
	ans int
}

func Test_Problem2275(t *testing.T) {
	qs := []question2275{
		{
			params{[]int{16, 17, 71, 62, 12, 24, 14}},
			4,
		},
		{
			params{[]int{8, 8}},
			2,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, largestCombination(p.candidates), ans)
	}
}
