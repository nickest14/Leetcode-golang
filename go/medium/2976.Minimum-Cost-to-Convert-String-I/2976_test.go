package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	source   string
	target   string
	original []byte
	changed  []byte
	cost     []int
}

type question2976 struct {
	params
	ans int64
}

func Test_Problem2976(t *testing.T) {
	qs := []question2976{
		{
			params{source: "abcd", target: "acbe", original: []byte{'a', 'b', 'c', 'c', 'e', 'd'}, changed: []byte{'b', 'c', 'b', 'e', 'b', 'e'}, cost: []int{2, 5, 5, 1, 2, 20}},
			28,
		},
		{
			params{source: "aaaa", target: "bbbb", original: []byte{'a', 'c'}, changed: []byte{'c', 'b'}, cost: []int{1, 2}},
			12,
		},
		{
			params{source: "abcd", target: "abce", original: []byte{'a'}, changed: []byte{'e'}, cost: []int{10000}},
			-1,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minimumCost(p.source, p.target, p.original, p.changed, p.cost), ans)
	}
}
