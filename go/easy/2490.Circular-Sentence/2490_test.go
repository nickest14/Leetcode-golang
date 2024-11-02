package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	sentence string
}

type question2490 struct {
	params
	ans bool
}

func Test_Problem2490(t *testing.T) {
	qs := []question2490{
		{
			params{"leetcode exercises sound delightful"},
			true,
		},
	}

	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, isCircularSentence(p.sentence), ans)
	}
}
