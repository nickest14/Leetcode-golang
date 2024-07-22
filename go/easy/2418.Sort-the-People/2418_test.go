package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	names   []string
	heights []int
}

type question2418 struct {
	params
	ans []string
}

func Test_Problem2418(t *testing.T) {
	qs := []question2418{
		{
			params{[]string{"Mary", "John", "Emma"}, []int{180, 165, 170}},
			[]string{"Mary", "Emma", "John"},
		},
	}

	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, sortPeople(p.names, p.heights), ans)
	}
}
