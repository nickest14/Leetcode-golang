package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	edges []int
	node1 int
	node2 int
}

type question2359 struct {
	params
	ans int
}

func Test_Problem2359(t *testing.T) {
	qs := []question2359{
		{
			params{edges: []int{2, 2, 3, -1}, node1: 0, node2: 1},
			2,
		},
		{
			params{edges: []int{1, 2, -1}, node1: 0, node2: 2},
			2,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, closestMeetingNode(p.edges, p.node1, p.node2), ans)
	}
}
