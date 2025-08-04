package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	fruits []int
}

type question904 struct {
	params
	ans int
}

func Test_Problem904(t *testing.T) {
	qs := []question904{
		{
			params{fruits: []int{1, 2, 1}},
			3,
		},
		{
			params{fruits: []int{0, 1, 2, 2}},
			3,
		},
		{
			params{fruits: []int{1, 2, 3, 2, 2}},
			4,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, totalFruit(p.fruits), ans)
	}
}
