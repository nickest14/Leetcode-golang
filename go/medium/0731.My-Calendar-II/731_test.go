package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	start int
	end   int
}

type question731 struct {
	params
	ans bool
}

func Test_Problem731(t *testing.T) {
	qs := []question731{
		{
			params{10, 20},
			true,
		},
		{
			params{50, 60},
			true,
		},
		{
			params{10, 40},
			true,
		},
		{
			params{5, 15},
			false,
		},
		{
			params{5, 10},
			true,
		},
		{
			params{25, 55},
			true,
		},
	}
	obj := Constructor()
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, obj.Book(p.start, p.end), ans)
	}
}
