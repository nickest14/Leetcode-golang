package leetcode

import (
	"fmt"
	"leetcode/go/structures"
	"testing"
)

type params struct {
	one     []int
	another []int
}

type question2 struct {
	params
	ans []int
}

func Test_Problem2(t *testing.T) {

	qs := []question2{
		{
			params{[]int{2, 4, 3}, []int{5, 6, 4}},
			[]int{7, 0, 8},
		},
		{
			params: params{[]int{}, []int{}},
			ans:    []int{},
		},
		{
			params{[]int{3}, []int{3}},
			[]int{6},
		},
		{
			params{[]int{1}, []int{9, 9, 9}},
			[]int{0, 0, 0, 1},
		},
		{
			params{[]int{9, 9, 9}, []int{1}},
			[]int{0, 0, 0, 1},
		},
		{
			params{[]int{3, 8, 3}, []int{7, 1}},
			[]int{0, 0, 4},
		},
	}

	for _, q := range qs {
		p := q.params
		fmt.Printf("[input]: %v        [output]: %v\n", p, structures.ListToInts(addTwoNumbers(structures.IntsToList(p.one), structures.IntsToList(p.another))))
	}
}
