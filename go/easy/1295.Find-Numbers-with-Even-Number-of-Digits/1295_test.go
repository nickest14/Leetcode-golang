package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question1295 struct {
	params
	ans int
}

func Test_Problem1295(t *testing.T) {
	qs := []question1295{
		{
			params{[]int{12, 345, 2, 6, 7896}},
			2,
		},
		{
			params{[]int{555, 901, 482, 1771}},
			1,
		},
	}

	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, findNumbers(p.nums), ans)
	}
}
