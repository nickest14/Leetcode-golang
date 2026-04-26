package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	grid [][]byte
}

type question1559 struct {
	params
	ans bool
}

func Test_Problem1559(t *testing.T) {
	qs := []question1559{
		{
			params{[][]byte{{'a', 'a', 'a', 'a'}, {'a', 'b', 'b', 'a'}, {'a', 'b', 'b', 'a'}, {'a', 'a', 'a', 'a'}}},
			true,
		},
		{
			params{[][]byte{{'c', 'c', 'c', 'a'}, {'c', 'd', 'c', 'c'}, {'c', 'c', 'e', 'c'}, {'f', 'c', 'c', 'c'}}},
			true,
		},
		{
			params{[][]byte{{'a', 'b', 'b'}, {'b', 'z', 'b'}, {'b', 'b', 'a'}}},
			false,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, containsCycle(p.grid), ans)
	}
}
