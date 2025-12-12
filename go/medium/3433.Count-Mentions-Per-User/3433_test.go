package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	numberOfUsers int
	events        [][]string
}

type question3433 struct {
	params
	ans []int
}

func Test_Problem3433(t *testing.T) {
	qs := []question3433{
		{
			params{numberOfUsers: 2, events: [][]string{{"MESSAGE", "10", "id1 id0"}, {"OFFLINE", "11", "0"}, {"MESSAGE", "71", "HERE"}}},
			[]int{2, 2},
		},
		{
			params{numberOfUsers: 2, events: [][]string{{"MESSAGE", "10", "id1 id0"}, {"OFFLINE", "11", "0"}, {"MESSAGE", "12", "ALL"}}},
			[]int{2, 2},
		},
		{
			params{numberOfUsers: 2, events: [][]string{{"OFFLINE", "10", "0"}, {"MESSAGE", "12", "HERE"}}},
			[]int{0, 1},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, countMentions(p.numberOfUsers, p.events), ans)
	}
}
