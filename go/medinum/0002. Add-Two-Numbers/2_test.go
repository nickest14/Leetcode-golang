package leetcode

import (
	"fmt"
	"leetcode/go/template"
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
		p, _ := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v\n", p, template.ListToInts(addTwoNumbers(template.IntsToList(p.one), template.IntsToList(p.another))))
	}
}
