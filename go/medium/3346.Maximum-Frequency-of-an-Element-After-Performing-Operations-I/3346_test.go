package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums          []int
	k             int
	numOperations int
}

type question3346 struct {
	params
	ans int
}

func Test_Problem3346(t *testing.T) {
	qs := []question3346{
		{
			params{nums: []int{1, 4, 5}, k: 1, numOperations: 2},
			2,
		},
		{
			params{nums: []int{5, 11, 20, 20}, k: 5, numOperations: 1},
			2,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maxFrequency(p.nums, p.k, p.numOperations), ans)
	}
}
