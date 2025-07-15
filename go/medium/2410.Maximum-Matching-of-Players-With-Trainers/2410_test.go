package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	players  []int
	trainers []int
}

type question2410 struct {
	params
	ans int
}

func Test_Problem2410(t *testing.T) {
	qs := []question2410{
		{
			params{players: []int{4, 7, 9}, trainers: []int{8, 2, 5, 8}},
			2,
		},
		{
			params{players: []int{1, 1, 1}, trainers: []int{10}},
			1,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, matchPlayersAndTrainers(p.players, p.trainers), ans)
	}
}
