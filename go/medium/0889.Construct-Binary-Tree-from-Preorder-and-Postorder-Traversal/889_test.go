package leetcode

import (
	"fmt"
	structures "leetcode/go/structures"
	"testing"
)

type params struct {
	preorder  []int
	postorder []int
}

type question889 struct {
	params
	ans []int
}

func Test_Problem889(t *testing.T) {
	qs := []question889{
		{
			params{preorder: []int{1, 2, 4, 5, 3, 6, 7}, postorder: []int{4, 5, 2, 6, 7, 3, 1}},
			[]int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			params{preorder: []int{1}, postorder: []int{1}},
			[]int{1},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, structures.TreeToints(constructFromPrePost(p.preorder, p.postorder)), ans)
	}
}
