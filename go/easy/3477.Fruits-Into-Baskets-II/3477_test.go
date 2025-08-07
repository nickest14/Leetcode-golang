package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	fruits  []int
	baskets []int
}

type question3477 struct {
	params
	ans int
}

func Test_Problem3477(t *testing.T) {
	qs := []question3477{
		{
			params{fruits: []int{4, 2, 5}, baskets: []int{3, 5, 4}},
			1,
		},
		{
			params{fruits: []int{3, 6, 1}, baskets: []int{6, 4, 7}},
			0,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, numOfUnplacedFruits(p.fruits, p.baskets), ans)
	}
}
