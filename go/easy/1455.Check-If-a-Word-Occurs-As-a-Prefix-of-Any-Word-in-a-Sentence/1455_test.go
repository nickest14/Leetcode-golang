package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	sentence   string
	searchWord string
}

type question1455 struct {
	params
	ans int
}

func Test_Problem1455(t *testing.T) {
	qs := []question1455{
		{
			params{"i love eating burger", "burg"},
			4,
		},
		{
			params{"this problem is an easy problem", "pro"},
			2,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, isPrefixOfWord(p.sentence, p.searchWord), ans)
	}
}
