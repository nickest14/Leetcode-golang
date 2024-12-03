package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s      string
	spaces []int
}

type question2109 struct {
	params
	ans string
}

func Test_Problem2109(t *testing.T) {
	qs := []question2109{
		{
			params{s: "LeetcodeHelpsMeLearn", spaces: []int{8, 13, 15}},
			"Leetcode Helps Me Learn",
		},
		{
			params{s: "icodeinpython", spaces: []int{1, 5, 7, 9}},
			"i code in py thon",
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, addSpaces(p.s, p.spaces), ans)
	}
}
