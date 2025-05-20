package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question3024 struct {
	params
	ans string
}

func Test_Problem3024(t *testing.T) {
	qs := []question3024{
		{
			params{nums: []int{3, 3, 3}},
			"equilateral",
		},
		{
			params{nums: []int{3, 4, 5}},
			"scalene",
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, triangleType(p.nums), ans)
	}
}
