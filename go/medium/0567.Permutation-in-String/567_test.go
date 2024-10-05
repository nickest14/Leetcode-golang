package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s1 string
	s2 string
}

type question567 struct {
	params
	ans bool
}

func Test_Problem567(t *testing.T) {
	qs := []question567{
		{
			params{s1: "ab", s2: "eidbaooo"},
			true,
		},
		{
			params{s1: "ab", s2: "eidboaoo"},
			false,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, checkInclusion(p.s1, p.s2), ans)
	}
}
