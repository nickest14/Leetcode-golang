package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	word string
	k    int
}

type question3085 struct {
	params
	ans int
}

func Test_Problem3085(t *testing.T) {
	qs := []question3085{
		{
			params{word: "aabcaba", k: 0},
			3,
		},
		{
			params{word: "dabdcbdcdcd", k: 2},
			2,
		},
		{
			params{word: "aaabaaa", k: 1},
			1,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minimumDeletions(p.word, p.k), ans)
	}
}
