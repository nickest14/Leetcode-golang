package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	str1 string
	str2 string
}

type question2528 struct {
	params
	ans bool
}

func Test_Problem2528(t *testing.T) {
	qs := []question2528{
		{
			params{str1: "abc", str2: "ad"},
			true,
		},
		{
			params{str1: "zc", str2: "ad"},
			true,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, canMakeSubsequence(p.str1, p.str2), ans)
	}
}
