package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	timePoints []string
}

type question539 struct {
	params
	ans int
}

func Test_Problem539(t *testing.T) {
	qs := []question539{
		{
			params{timePoints: []string{"23:59", "00:00"}},
			1,
		},
		{
			params{timePoints: []string{"00:00", "23:00", "22:30"}},
			30,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, findMinDifference(p.timePoints), ans)
	}
}
