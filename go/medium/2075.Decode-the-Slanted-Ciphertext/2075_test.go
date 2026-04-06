package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	encodedText string
	rows        int
}

type question2075 struct {
	params
	ans string
}

func Test_Problem2075(t *testing.T) {
	qs := []question2075{
		{
			params{encodedText: "ch   ie   pr", rows: 3},
			"cipher",
		},
		{
			params{encodedText: "iveo    eed   l te   olc", rows: 4},
			"i love leetcode",
		},
		{
			params{encodedText: "coding", rows: 1},
			"coding",
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, decodeCiphertext(p.encodedText, p.rows), ans)
	}
}
