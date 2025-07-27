package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s string
	x int
	y int
}

type question1717 struct {
	params
	ans int
}

func Test_Problem1717(t *testing.T) {
	qs := []question1717{
		{
			params{s: "cdbcbbaaabab", x: 4, y: 5},
			19,
		},
		{
			params{s: "aabbaaxybbaabb", x: 5, y: 4},
			20,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maximumGain(p.s, p.x, p.y), ans)
	}
}
