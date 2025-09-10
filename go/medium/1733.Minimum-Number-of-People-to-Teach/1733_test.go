package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n           int
	languages   [][]int
	friendships [][]int
}

type question1733 struct {
	params
	ans int
}

func Test_Problem1733(t *testing.T) {
	qs := []question1733{
		{
			params{n: 2, languages: [][]int{{1}, {2}, {1, 2}}, friendships: [][]int{{1, 2}, {1, 3}, {2, 3}}},
			1,
		},
		{
			params{n: 3, languages: [][]int{{2}, {1, 3}, {1, 2}, {3}}, friendships: [][]int{{1, 4}, {1, 2}, {3, 4}, {2, 3}}},
			2,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minimumTeachings(p.n, p.languages, p.friendships), ans)
	}
}
