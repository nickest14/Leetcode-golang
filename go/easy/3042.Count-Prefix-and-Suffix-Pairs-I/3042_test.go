package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	words []string
}

type question3042 struct {
	params
	ans int
}

func Test_Problem3042(t *testing.T) {
	qs := []question3042{
		{
			params{[]string{"a", "aba", "ababa", "aa"}},
			4,
		},
		{
			params{[]string{"pa", "papa", "ma", "mama"}},
			2,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, countPrefixSuffixPairs(p.words), ans)
	}
}
