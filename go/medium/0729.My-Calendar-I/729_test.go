package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	start int
	end   int
}

type question729 struct {
	params
	ans bool
}

func Test_Problem729(t *testing.T) {
	qs := []question729{
		{
			params{10, 20},
			true,
		},
		{
			params{15, 25},
			false,
		},
		{
			params{20, 30},
			true,
		},
	}
	obj := Constructor()
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, obj.Book(p.start, p.end), ans)
	}
}
