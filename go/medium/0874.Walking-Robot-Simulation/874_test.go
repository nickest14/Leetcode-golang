package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	commands  []int
	obstacles [][]int
}

type question874 struct {
	params
	ans int
}

func Test_Problem874(t *testing.T) {
	qs := []question874{
		{
			params{commands: []int{4, -1, 3}, obstacles: [][]int{}},
			25,
		},
		{
			params{commands: []int{4, -1, 4, -2, 4}, obstacles: [][]int{{2, 4}}},
			65,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, robotSim(p.commands, p.obstacles), ans)
	}
}
