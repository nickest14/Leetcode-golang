package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	version1 string
	version2 string
}

type question165 struct {
	params
	ans int
}

func Test_Problem165(t *testing.T) {

	qs := []question165{
		{
			params{version1: "1.2", version2: "1.10"},
			-1,
		},
		{
			params{version1: "1.01", version2: "1.001"},
			0,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v, [answer]: %v\n", p, compareVersion(p.version1, p.version2), ans)
	}
}
