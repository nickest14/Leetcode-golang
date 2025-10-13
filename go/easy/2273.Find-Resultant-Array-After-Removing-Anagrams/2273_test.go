package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	words []string
}

type question2273 struct {
	params
	ans []string
}

func Test_Problem2273(t *testing.T) {
	qs := []question2273{
		{
			params{[]string{"abba", "baba", "bbaa", "cd", "cd"}},
			[]string{"abba", "cd"},
		},
		{
			params{[]string{"a", "b", "c", "d", "e"}},
			[]string{"a", "b", "c", "d", "e"},
		},
	}

	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, removeAnagrams(p.words), ans)
	}
}
