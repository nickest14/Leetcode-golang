package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	happiness []int
	k         int
}

type question3075 struct {
	params
	ans int64
}

func Test_Problem3075(t *testing.T) {
	qs := []question3075{
		{
			params{happiness: []int{1, 2, 3}, k: 2},
			4,
		},
		{
			params{happiness: []int{1, 1, 1, 1}, k: 2},
			1,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maximumHappinessSum(p.happiness, p.k), ans)
	}
}
