package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	gifts []int
	k     int
}

type question2558 struct {
	params
	ans int64
}

func Test_Problem2558(t *testing.T) {
	qs := []question2558{
		{
			params{gifts: []int{25, 64, 9, 4, 100}, k: 4},
			29,
		},
		{
			params{gifts: []int{1, 1, 1, 1}, k: 4},
			4,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, pickGifts(p.gifts, p.k), ans)
	}
}
