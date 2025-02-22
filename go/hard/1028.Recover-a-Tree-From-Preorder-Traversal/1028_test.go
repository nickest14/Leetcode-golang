package leetcode

import (
	"fmt"
	structures "leetcode/go/structures"
	"testing"
)

type params struct {
	traversal string
}

type question1028 struct {
	params
	ans []int
}

func Test_Problem1028(t *testing.T) {
	qs := []question1028{
		{
			params{"1-2--3--4-5--6--7"},
			[]int{1, 2, 5, 3, 4, 6, 7},
		},
		{
			params{"1-2--3---4-5--6---7"},
			[]int{1, 2, 5, 3, structures.NULL, 6, structures.NULL, 4, structures.NULL, 7},
		},
		{
			params{"1-401--349---90--88"},
			[]int{1, 401, structures.NULL, 349, 88, 90},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, structures.TreeToints(recoverFromPreorder(p.traversal)), ans)
	}
}
