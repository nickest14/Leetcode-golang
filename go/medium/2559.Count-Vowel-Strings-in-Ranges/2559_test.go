package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	words   []string
	queries [][]int
}

type question2559 struct {
	params
	ans []int
}

func Test_Problem2559(t *testing.T) {
	qs := []question2559{
		{
			params{words: []string{"aba", "bcb", "ece", "aa", "e"}, queries: [][]int{{0, 2}, {1, 4}, {1, 1}}},
			[]int{2, 3, 0},
		},
		{
			params{words: []string{"a", "e", "i"}, queries: [][]int{{0, 2}, {0, 1}, {2, 2}}},
			[]int{3, 2, 1},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, vowelStrings(p.words, p.queries), ans)
	}
}
