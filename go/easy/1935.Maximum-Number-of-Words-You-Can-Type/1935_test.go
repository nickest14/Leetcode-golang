package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	text          string
	brokenLetters string
}

type question935 struct {
	params
	ans int
}

func Test_Problem935(t *testing.T) {
	qs := []question935{
		{
			params{text: "hello world", brokenLetters: "ad"},
			1,
		},
		{
			params{text: "leet code", brokenLetters: "lt"},
			1,
		},
		{
			params{text: "leet code", brokenLetters: "e"},
			0,
		},
	}

	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, canBeTypedWords(p.text, p.brokenLetters), ans)
	}
}
