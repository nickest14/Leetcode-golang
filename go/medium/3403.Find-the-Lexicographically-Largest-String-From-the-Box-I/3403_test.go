package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	word       string
	numFriends int
}

type question3403 struct {
	params
	ans string
}

func Test_Problem3403(t *testing.T) {
	qs := []question3403{
		// {
		// 	params{word: "dbca", numFriends: 2},
		// 	"dbc",
		// },
		// {
		// 	params{word: "gggg", numFriends: 4},
		// 	"g",
		// },
		// {
		// 	params{word: "aann", numFriends: 2},
		// 	"nn",
		// },
		{
			params{word: "gh", numFriends: 1},
			"gh",
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, answerString(p.word, p.numFriends), ans)
	}
}
