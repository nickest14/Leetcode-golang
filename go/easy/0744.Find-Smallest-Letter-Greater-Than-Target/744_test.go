package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	letters []byte
	target  byte
}

type question744 struct {
	params
	ans byte
}

func Test_Problem744(t *testing.T) {
	qs := []question744{
		{
			params{letters: []byte{'c', 'f', 'j'}, target: 'a'},
			'c',
		},
		{
			params{letters: []byte{'c', 'f', 'j'}, target: 'c'},
			'f',
		},
		{
			params{letters: []byte{'x', 'x', 'y', 'y'}, target: 'z'},
			'x',
		},
	}

	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, nextGreatestLetter(p.letters, p.target), ans)
	}
}
