package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	word string
}

type question3136 struct {
	params
	ans bool
}

func Test_Problem3136(t *testing.T) {
	qs := []question3136{
		{
			params{"234Adas"},
			true,
		},
		{
			params{"b3"},
			false,
		},
		{
			params{"a3$e"},
			false,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, isValid(p.word), ans)
	}
}
