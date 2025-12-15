package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	prices []int
}

type question2110 struct {
	params
	ans int64
}

func Test_Problem2110(t *testing.T) {
	qs := []question2110{
		{
			params{prices: []int{3, 2, 1, 4}},
			7,
		},
		{
			params{prices: []int{8, 6, 7, 7}},
			4,
		},
		{
			params{prices: []int{1}},
			1,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, getDescentPeriods(p.prices), ans)
	}
}
