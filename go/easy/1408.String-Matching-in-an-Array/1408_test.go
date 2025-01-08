package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	words []string
}

type question1408 struct {
	params
	ans []string
}

func Test_Problem1408(t *testing.T) {
	qs := []question1408{
		{
			params{[]string{"mass", "as", "hero", "superhero"}},
			[]string{"as", "hero"},
		},
		{
			params{[]string{"leetcode", "et", "code"}},
			[]string{"et", "code"},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, stringMatching(p.words), ans)
	}
}
