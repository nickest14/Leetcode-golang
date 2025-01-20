package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	arr []int
	mat [][]int
}

type question2661 struct {
	params
	ans int
}

func Test_Problem2661(t *testing.T) {
	qs := []question2661{
		{
			params{arr: []int{1, 3, 4, 2}, mat: [][]int{{1, 4}, {2, 3}}},
			2,
		},
		{
			params{arr: []int{2, 8, 7, 4, 1, 3, 5, 6, 9}, mat: [][]int{{3, 2, 5}, {1, 4, 6}, {8, 7, 9}}},
			3,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, firstCompleteIndex(p.arr, p.mat), ans)
	}
}
