package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	complexity []int
}

type question3577 struct {
	params
	ans int
}

func Test_Problem3577(t *testing.T) {
	qs := []question3577{
		{
			params{complexity: []int{1, 2, 3}},
			2,
		},
		{
			params{complexity: []int{3, 3, 3, 4, 4, 4}},
			0,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, countPermutations(p.complexity), ans)
	}
}
