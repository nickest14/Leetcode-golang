package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	apple    []int
	capacity []int
}

type question3074 struct {
	params
	ans int
}

func Test_Problem3074(t *testing.T) {
	qs := []question3074{
		{
			params{apple: []int{1, 3, 2}, capacity: []int{4, 3, 1, 5, 2}},
			2,
		},
		{
			params{apple: []int{5, 5, 5}, capacity: []int{2, 4, 2, 7}},
			4,
		},
	}

	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minimumBoxes(p.apple, p.capacity), ans)
	}
}
