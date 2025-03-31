package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	s string
}

type question763 struct {
	params
	ans []int
}

func Test_Problem763(t *testing.T) {
	qs := []question763{
		// {
		// 	params{"ababcbacadefegdehijhklij"},
		// 	[]int{9, 7, 8},
		// },
		{
			params{"eccbbbbdec"},
			[]int{10},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, partitionLabels(p.s), ans)
	}
}
