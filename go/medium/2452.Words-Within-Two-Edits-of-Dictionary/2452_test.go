package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	queries    []string
	dictionary []string
}

type question2452 struct {
	params
	ans []string
}

func Test_Problem2452(t *testing.T) {
	qs := []question2452{
		{
			params{queries: []string{"word", "note", "ants", "wood"}, dictionary: []string{"wood", "joke", "moat"}},
			[]string{"word", "note", "wood"},
		},
		{
			params{queries: []string{"yes"}, dictionary: []string{"not"}},
			[]string{},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, twoEditWords(p.queries, p.dictionary), ans)
	}
}
