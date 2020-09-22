package leetcode

import (
	"fmt"
	"testing"
)

type question3 struct {
	para
	ans
}

type para struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	length int
}

func Test_Problem3(t *testing.T) {

	qs := []question3{
		{
			para{"abcdeabc"},
			ans{5},
		},
		{
			para{"bbb"},
			ans{1},
		},
		{
			para{"quuvuula"},
			ans{3},
		},
		{
			para{""},
			ans{0},
		},
		{
			para{"abcdefg"},
			ans{7},
		},
	}
	for _, q := range qs {
		p, a := q.para, q.ans
		fmt.Printf("[input]: %v        [output]: %v, [answer]: %v\n", p, lengthOfLongestSubstring(p.s), a.length)
	}
}
