package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	words  []string
	target string
}

type question1639 struct {
	params
	ans int
}

func Test_Problem1639(t *testing.T) {
	qs := []question1639{
		{
			params{words: []string{"acca", "bbbb", "caca"}, target: "aba"},
			6,
		},
		{
			params{words: []string{"abba", "baab"}, target: "bab"},
			4,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, numWays(p.words, p.target), ans)
	}
}
