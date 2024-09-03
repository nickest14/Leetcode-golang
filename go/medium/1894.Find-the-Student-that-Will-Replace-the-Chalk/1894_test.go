package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	chalk []int
	k     int
}

type question1894 struct {
	params
	ans int
}

func Test_Problem1894(t *testing.T) {
	qs := []question1894{
		{
			params{chalk: []int{5, 1, 5}, k: 22},
			22,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, chalkReplacer(p.chalk, p.k), ans)
	}
}
