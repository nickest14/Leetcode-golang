package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	arr []int
}

type question769 struct {
	params
	ans int
}

func Test_Problem769(t *testing.T) {
	qs := []question769{
		{
			params{arr: []int{4, 3, 2, 1, 0}},
			1,
		},
		{
			params{arr: []int{1, 0, 2, 3, 4}},
			4,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maxChunksToSorted(p.arr), ans)
	}
}
