package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	weights []int
	k       int
}

type question2551 struct {
	params
	ans int64
}

func Test_Problem2551(t *testing.T) {
	qs := []question2551{
		{
			params{weights: []int{1, 3, 5, 1}, k: 2},
			4,
		},
		{
			params{weights: []int{1, 3}, k: 2},
			0,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, putMarbles(p.weights, p.k), ans)
	}
}
