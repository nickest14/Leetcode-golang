package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	n           int
	meetings    [][]int
	firstPerson int
}

type question2092 struct {
	params
	ans []int
}

func Test_Problem2092(t *testing.T) {
	qs := []question2092{
		{
			params{n: 6, meetings: [][]int{{1, 2, 5}, {2, 3, 8}, {1, 5, 10}}, firstPerson: 1},
			[]int{0, 1, 2, 3, 5},
		},
		{
			params{n: 4, meetings: [][]int{{3, 1, 3}, {1, 2, 2}, {0, 3, 3}}, firstPerson: 3},
			[]int{0, 1, 3},
		},
		{
			params{n: 5, meetings: [][]int{{3, 4, 2}, {1, 2, 1}, {2, 3, 1}}, firstPerson: 1},
			[]int{0, 1, 2, 3, 4},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, findAllPeople(p.n, p.meetings, p.firstPerson), ans)
	}
}
