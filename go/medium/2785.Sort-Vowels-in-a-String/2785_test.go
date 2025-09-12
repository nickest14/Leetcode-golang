package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s string
}

type question2785 struct {
	params
	ans string
}

func Test_Problem2785(t *testing.T) {
	qs := []question2785{
		{
			params{"lEetcOde"},
			"lEOtcede",
		},
		{
			params{"lYmpH"},
			"lYmpH",
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, sortVowels(p.s), ans)
	}
}
