package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	bottom  string
	allowed []string
}

type question889 struct {
	params
	ans bool
}

func Test_Problem889(t *testing.T) {
	qs := []question889{
		{
			params{bottom: "BCD", allowed: []string{"BCC", "CDE", "CEA", "FFF"}},
			true,
		},
		{
			params{bottom: "AAAA", allowed: []string{"AAB", "AAC", "BCD", "BBE", "DEF"}},
			false,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, pyramidTransition(p.bottom, p.allowed), ans)
	}
}
