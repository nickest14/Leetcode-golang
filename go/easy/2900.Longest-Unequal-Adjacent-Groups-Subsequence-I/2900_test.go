package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	words  []string
	groups []int
}

type question2900 struct {
	params
	ans []string
}

func Test_Problem2900(t *testing.T) {
	qs := []question2900{
		{
			params{words: []string{"e", "a", "b"}, groups: []int{0, 0, 1}},
			[]string{"e", "b"},
		},
		{
			params{words: []string{"a", "b", "c", "d"}, groups: []int{1, 0, 1, 1}},
			[]string{"a", "b", "c"},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, getLongestSubsequence(p.words, p.groups), ans)
	}
}
