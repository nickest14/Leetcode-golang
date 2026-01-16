package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n     int
	m     int
	hBars []int
	vBars []int
}
type question2943 struct {
	params
	ans int
}

func Test_Problem2943(t *testing.T) {
	qs := []question2943{
		{
			params{n: 2, m: 1, hBars: []int{2, 3}, vBars: []int{2}},
			4,
		},
		{
			params{n: 1, m: 1, hBars: []int{2}, vBars: []int{2}},
			4,
		},
		{
			params{n: 2, m: 3, hBars: []int{2, 3}, vBars: []int{2, 4}},
			4,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maximizeSquareHoleArea(p.n, p.m, p.hBars, p.vBars), ans)
	}
}
