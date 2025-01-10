package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	words1 []string
	words2 []string
}

type question916 struct {
	params
	ans []string
}

func Test_Problem916(t *testing.T) {
	qs := []question916{
		{
			params{words1: []string{"amazon", "apple", "facebook", "google", "leetcode"}, words2: []string{"e", "o"}},
			[]string{"facebook", "google", "leetcode"},
		},
		{
			params{words1: []string{"amazon", "apple", "facebook", "google", "leetcode"}, words2: []string{"l", "e"}},
			[]string{"apple", "google", "leetcode"},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, wordSubsets(p.words1, p.words2), ans)
	}
}
