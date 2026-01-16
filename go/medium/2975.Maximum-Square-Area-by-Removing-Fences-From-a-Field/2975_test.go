package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n       int
	m       int
	hFences []int
	vFences []int
}
type question2975 struct {
	params
	ans int
}

func Test_Problem2975(t *testing.T) {
	qs := []question2975{
		{
			params{n: 4, m: 3, hFences: []int{2, 3}, vFences: []int{2}},
			4,
		},
		{
			params{n: 6, m: 7, hFences: []int{2}, vFences: []int{4}},
			-1,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maximizeSquareArea(p.n, p.m, p.hFences, p.vFences), ans)
	}
}
