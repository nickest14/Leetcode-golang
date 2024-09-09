package leetcode

import (
	"fmt"
	structures "leetcode/go/structures"
	"testing"
)

type params struct {
	m    int
	n    int
	head *structures.ListNode
}

type question2326 struct {
	params
	ans [][]int
}

func Test_Problem2326(t *testing.T) {
	qs := []question2326{
		{
			params{3, 5, structures.IntsToList([]int{3, 0, 2, 6, 8, 1, 7, 9, 4, 2, 5, 5, 0})},
			[][]int{{3, 0, 2, 6, 8}, {5, 0, -1, -1, 1}, {5, 2, 4, 9, 7}},
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, spiralMatrix(p.m, p.n, p.head), ans)
	}
}
