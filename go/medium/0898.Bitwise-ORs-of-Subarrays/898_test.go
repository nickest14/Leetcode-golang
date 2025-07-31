package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	arr []int
}

type question898 struct {
	params
	ans int
}

func Test_Problem898(t *testing.T) {
	qs := []question898{
		{
			params{arr: []int{0}},
			1,
		},
		{
			params{arr: []int{1, 1, 2}},
			3,
		},
		{
			params{arr: []int{1, 2, 4}},
			6,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, subarrayBitwiseORs(p.arr), ans)
	}
}
