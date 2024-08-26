package leetcode

import (
	"fmt"
	"leetcode/go/structures"
	"testing"
)

type params struct {
	nums []int
}

type question590 struct {
	params
	ans []int
}

func Test_Problem590(t *testing.T) {

	qs := []question590{
		{
			params{[]int{1, structures.NULL, 3, 2, 4, structures.NULL, 5, 6}},
			[]int{5, 6, 3, 2, 4, 1},
		},
	}

	for _, q := range qs {
		p := q.params
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, postorder(structures.IntsToNode(p.nums)), q.ans)
	}
}
