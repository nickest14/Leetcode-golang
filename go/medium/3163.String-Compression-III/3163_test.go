package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	word string
}

type question3163 struct {
	params
	ans string
}

func Test_Problem3163(t *testing.T) {
	qs := []question3163{
		{
			params{"aaaaaaaaaaaaaabb"},
			"9a5a2b",
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, compressedString(p.word), ans)
	}
}
