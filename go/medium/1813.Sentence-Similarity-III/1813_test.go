package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	sentence1 string
	sentence2 string
}

type question1813 struct {
	params
	ans bool
}

func Test_Problem1813(t *testing.T) {
	qs := []question1813{
		{
			params{sentence1: "My name is Haley", sentence2: "My Haley"},
			true,
		},
		{
			params{sentence1: "of", sentence2: "A lot of words"},
			false,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, areSentencesSimilar(p.sentence1, p.sentence2), ans)
	}
}
