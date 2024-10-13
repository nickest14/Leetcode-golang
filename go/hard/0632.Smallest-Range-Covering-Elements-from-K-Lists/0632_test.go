package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums [][]int
}

type question632 struct {
	params
	ans []int
}

func Test_Problem632(t *testing.T) {
	qs := []question632{
		{
			params{[][]int{{4, 10, 15, 24, 26}, {0, 9, 12, 20}, {5, 18, 22, 30}}},
			[]int{20, 24},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, smallestRange(p.nums), ans)
	}
}
