package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []string
}

type question1980 struct {
	params
	ans string
}

func Test_Problem1980(t *testing.T) {
	qs := []question1980{
		{
			params{[]string{"01", "10"}},
			"11",
		},
		{
			params{[]string{"00", "01"}},
			"10",
		},
		{
			params{[]string{"111", "011", "001"}},
			"000",
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, findDifferentBinaryString(p.nums), ans)
	}
}
