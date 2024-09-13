package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	allowed string
	words   []string
}

type question1684 struct {
	params
	ans int
}

func Test_Problem1684(t *testing.T) {
	qs := []question1684{
		{
			params{"ab", []string{"ad", "bd", "aaab", "baa", "badab"}},
			2,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, countConsistentStrings(p.allowed, p.words), ans)
	}
}
