package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	words []string
	pref  string
}

type question2185 struct {
	params
	ans int
}

func Test_Problem2185(t *testing.T) {
	qs := []question2185{
		{
			params{words: []string{"pay", "attention", "practice", "attend"}, pref: "at"},
			2,
		},
		{
			params{words: []string{"leetcode", "win", "loops", "success"}, pref: "code"},
			0,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, prefixCount(p.words, p.pref), ans)
	}
}
