package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	code []int
	k    int
}

type question1652 struct {
	params
	ans []int
}

func Test_Problem1652(t *testing.T) {
	qs := []question1652{
		// {
		// 	params{code: []int{5, 7, 1, 4}, k: 3},
		// 	[]int{12, 10, 16, 13},
		// },
		// {
		// 	params{code: []int{2, 4, 9, 3}, k: -2},
		// 	[]int{12, 5, 6, 13},
		// },
		{
			params{code: []int{5, 2, 2, 3, 1}, k: 3},
			[]int{7, 6, 9, 8, 9},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, decrypt(p.code, p.k), ans)
	}
}
