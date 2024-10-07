package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s string
}

type question2696 struct {
	params
	ans int
}

func Test_Problem2696(t *testing.T) {
	qs := []question2696{
		{
			params{"ABFCACDB"},
			2,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minLength(p.s), ans)
	}
}
