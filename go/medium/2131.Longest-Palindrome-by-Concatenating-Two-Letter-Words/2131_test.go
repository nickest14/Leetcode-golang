package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	words []string
}

type question2131 struct {
	params
	ans int
}

func Test_Problem2131(t *testing.T) {
	qs := []question2131{
		{
			params{[]string{"lc", "cl", "gg"}},
			6,
		},
		{
			params{[]string{"ab", "ty", "yt", "lc", "cl", "ab"}},
			8,
		},
		{
			params{[]string{"cc", "ll", "xx"}},
			2,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, longestPalindrome(p.words), ans)
	}
}
