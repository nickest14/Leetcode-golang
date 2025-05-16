package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	words  []string
	groups []int
}

type question2901 struct {
	params
	ans []string
}

func Test_Problem2901(t *testing.T) {
	qs := []question2901{
		{
			params{words: []string{"bab", "dab", "cab"}, groups: []int{1, 2, 2}},
			[]string{"bab", "dab"},
		},
		{
			params{words: []string{"a", "b", "c", "d"}, groups: []int{1, 2, 3, 4}},
			[]string{"a", "b", "c", "d"},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, getWordsInLongestSubsequence(p.words, p.groups), ans)
	}
}
