package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	wordlist []string
	queries  []string
}

type question966 struct {
	params
	ans []string
}

func Test_Problem966(t *testing.T) {
	qs := []question966{
		{
			params{wordlist: []string{"KiTe", "kite", "hare", "Hare"}, queries: []string{"kite", "Kite", "KiTe", "Hare", "HARE", "Hear", "hear", "keti", "keet", "keto"}},
			[]string{"kite", "KiTe", "KiTe", "Hare", "hare", "", "", "KiTe", "", "KiTe"},
		},
		{
			params{wordlist: []string{"yellow"}, queries: []string{"YellOw"}},
			[]string{"yellow"},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, spellchecker(p.wordlist, p.queries), ans)
	}
}
