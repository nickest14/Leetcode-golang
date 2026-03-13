package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	mountainHeight int
	workerTimes    []int
}

type question3296 struct {
	params
	ans int64
}

func Test_Problem3296(t *testing.T) {
	qs := []question3296{
		{
			params{mountainHeight: 4, workerTimes: []int{2, 1, 1}},
			3,
		},
		{
			params{mountainHeight: 10, workerTimes: []int{3, 2, 2, 4}},
			12,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minNumberOfSeconds(p.mountainHeight, p.workerTimes), ans)
	}
}
