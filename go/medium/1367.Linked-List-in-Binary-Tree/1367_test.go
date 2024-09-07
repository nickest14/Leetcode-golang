package leetcode

import (
	"fmt"
	structures "leetcode/go/structures"
	"testing"
)

type params struct {
	head *structures.ListNode
	root *structures.TreeNode
}

type question1367 struct {
	params
	ans bool
}

func Test_Problem1367(t *testing.T) {
	qs := []question1367{
		{
			params{
				structures.IntsToList([]int{4, 2, 8}),
				structures.IntsToTreeNode([]int{1, 4, 4, structures.NULL, 2, 2, structures.NULL, 1, structures.NULL, 6, 8, structures.NULL, structures.NULL, structures.NULL, structures.NULL, 1, 3})},
			true,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, isSubPath(p.head, p.root), ans)
	}
}
