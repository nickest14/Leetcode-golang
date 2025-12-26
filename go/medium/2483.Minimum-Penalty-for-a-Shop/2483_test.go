package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	customers string
}

type question2483 struct {
	params
	ans int
}

func Test_Problem2483(t *testing.T) {
	qs := []question2483{
		{
			params{customers: "YYNY"},
			2,
		},
		{
			params{customers: "NNNNN"},
			0,
		},
		{
			params{customers: "YYYY"},
			4,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, bestClosingTime(p.customers), ans)
	}
}
