package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	operations []string
}

type question2011 struct {
	params
	ans int
}

func Test_Problem2011(t *testing.T) {
	qs := []question2011{
		{
			params{operations: []string{"--X", "X++", "X++"}},
			1,
		},
		{
			params{operations: []string{"++X", "++X", "X++"}},
			3,
		},
		{
			params{operations: []string{"X++", "++X", "--X", "X--"}},
			0,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, finalValueAfterOperations(p.operations), ans)
	}
}
