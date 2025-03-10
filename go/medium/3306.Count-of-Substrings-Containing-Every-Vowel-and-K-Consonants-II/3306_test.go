package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	word string
	k    int
}

type question3306 struct {
	params
	ans int
}

func Test_Problem3306(t *testing.T) {
	qs := []question3306{
		{
			params{word: "aeioqq", k: 1},
			0,
		},
		{
			params{word: "aeiou", k: 0},
			1,
		},
		{
			params{word: "ieaouqqieaouqq", k: 1},
			3,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, countOfSubstrings(p.word, p.k), ans)
	}
}
