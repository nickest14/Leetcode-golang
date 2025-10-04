package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	numBottles  int
	numExchange int
}

type question1518 struct {
	params
	ans int
}

func Test_Problem1518(t *testing.T) {
	qs := []question1518{
		{
			params{numBottles: 9, numExchange: 3},
			13,
		},
		{
			params{numBottles: 15, numExchange: 4},
			19,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, numWaterBottles(p.numBottles, p.numExchange), ans)
	}
}
