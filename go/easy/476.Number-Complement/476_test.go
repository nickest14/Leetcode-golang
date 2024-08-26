package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	num int
}

type question1380 struct {
	params
	ans int
}

func Test_Problem1380(t *testing.T) {
	qs := []question1380{
		{
			params{5},
			2,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, findComplement(p.num), ans)
	}
}
