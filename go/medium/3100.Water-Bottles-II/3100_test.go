package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	numBottles  int
	numExchange int
}

type question3100 struct {
	params
	ans int
}

func Test_Problem3100(t *testing.T) {
	qs := []question3100{
		{
			params{numBottles: 13, numExchange: 6},
			15,
		},
		{
			params{numBottles: 10, numExchange: 3},
			13,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maxBottlesDrunk(p.numBottles, p.numExchange), ans)
	}
}
