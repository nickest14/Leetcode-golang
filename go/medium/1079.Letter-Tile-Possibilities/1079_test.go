package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	tiles string
}

type question1079 struct {
	params
	ans int
}

func Test_Problem1079(t *testing.T) {
	qs := []question1079{
		{
			params{"AAB"},
			8,
		},
		{
			params{"AAABBC"},
			188,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, numTilePossibilities(p.tiles), ans)
	}
}
