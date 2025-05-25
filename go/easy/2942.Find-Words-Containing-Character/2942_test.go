package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	words []string
	x     byte
}

type question2942 struct {
	params
	ans []int
}

func Test_Problem2942(t *testing.T) {
	qs := []question2942{
		{
			params{words: []string{"leet", "code"}, x: 'e'},
			[]int{0, 1},
		},
		{
			params{words: []string{"abc", "bcd", "aaaa", "cbc"}, x: 'a'},
			[]int{0, 2},
		},
		{
			params{words: []string{"abc", "bcd", "aaaa", "cbc"}, x: 'z'},
			[]int{},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, findWordsContaining(p.words, p.x), ans)
	}
}
