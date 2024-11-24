package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	box [][]byte
}

type question1861 struct {
	params
	ans [][]byte
}

func Test_Problem1861(t *testing.T) {
	qs := []question1861{
		{
			params{[][]byte{{'#', '#', '*', '.', '*', '.'}, {'#', '#', '#', '*', '.', '.'}, {'#', '#', '#', '.', '#', '.'}}},
			[][]byte{{'.', '#', '#'}, {'.', '#', '#'}, {'#', '#', '*'}, {'#', '*', '.'}, {'#', '.', '*'}, {'#', '.', '.'}},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, rotateTheBox(p.box), ans)
	}
}
