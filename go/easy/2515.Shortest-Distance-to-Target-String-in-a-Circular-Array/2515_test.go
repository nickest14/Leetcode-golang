package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	words      []string
	target     string
	startIndex int
}

type question2210 struct {
	params
	ans int
}

func Test_Problem2210(t *testing.T) {
	qs := []question2210{
		{
			params{words: []string{"hello", "i", "am", "leetcode", "hello"}, target: "hello", startIndex: 1},
			1,
		},
		{
			params{words: []string{"a", "b", "leetcode"}, target: "leetcode", startIndex: 0},
			1,
		},
		{
			params{words: []string{"i", "eat", "leetcode"}, target: "ate", startIndex: 0},
			-1,
		},
	}

	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, closestTarget(p.words, p.target, p.startIndex), ans)
	}
}
