package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	energy []int
	k      int
}

type question3147 struct {
	params
	ans int
}

func Test_Problem3147(t *testing.T) {
	qs := []question3147{
		{
			params{energy: []int{5, 2, -10, -5, 1}, k: 3},
			3,
		},
		{
			params{energy: []int{-2, -3, -1}, k: 2},
			-1,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maximumEnergy(p.energy, p.k), ans)
	}
}
