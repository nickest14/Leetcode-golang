package leetcode

import (
	"fmt"
	structures "leetcode/go/structures"
	"testing"
)

type params struct {
	values1 []int
	values2 []int
}

type question951 struct {
	params
	ans bool
}

func Test_Problem951(t *testing.T) {
	qs := []question951{
		{
			params{
				values1: []int{1, 2, 3, 4, 5, 6, structures.NULL, structures.NULL, structures.NULL, 7, 8},
				values2: []int{1, 3, 2, structures.NULL, 6, 4, 5, structures.NULL, structures.NULL, structures.NULL, structures.NULL, 8, 7},
			},
			true,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		root1 := structures.IntsToTreeNode(p.values1)
		root2 := structures.IntsToTreeNode(p.values2)
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, flipEquiv(root1, root2), ans)
	}
}
