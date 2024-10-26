package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	folder []string
}

type question1310 struct {
	params
	ans []string
}

func Test_Problem1310(t *testing.T) {
	qs := []question1310{
		{
			params{[]string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}},
			[]string{"/a", "/c/d", "/c/f"},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, removeSubfolders(p.folder), ans)
	}
}
