package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	colors []int
	k      int
}

type question3208 struct {
	params
	ans int
}

func Test_Problem3208(t *testing.T) {
	qs := []question3208{
		{
			params{colors: []int{0, 1, 0, 1, 0}, k: 3},
			3,
		},
		{
			params{colors: []int{0, 1, 0, 0, 1, 0, 15}, k: 6},
			2,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, numberOfAlternatingGroups(p.colors, p.k), ans)
	}
}
