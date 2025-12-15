package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	code         []string
	businessLine []string
	isActive     []bool
}

type question3606 struct {
	params
	ans []string
}

func Test_Problem3606(t *testing.T) {
	qs := []question3606{
		{
			params{code: []string{"SAVE20", "", "PHARMA5", "SAVE@20"}, businessLine: []string{"restaurant", "grocery", "pharmacy", "restaurant"}, isActive: []bool{true, true, true, true}},
			[]string{"PHARMA5", "SAVE20"},
		},
		{
			params{code: []string{"GROCERY15", "ELECTRONICS_50", "DISCOUNT10"}, businessLine: []string{"grocery", "electronics", "invalid"}, isActive: []bool{false, true, true}},
			[]string{"ELECTRONICS_50"},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, validateCoupons(p.code, p.businessLine, p.isActive), ans)
	}
}
