package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	status         []int
	candies        []int
	keys           [][]int
	containedBoxes [][]int
	initialBoxes   []int
}

type question1298 struct {
	params
	ans int
}

func Test_Problem1298(t *testing.T) {
	qs := []question1298{
		{
			params{status: []int{1, 0, 1, 0}, candies: []int{7, 5, 4, 100}, keys: [][]int{{}, {}, {1}, {}}, containedBoxes: [][]int{{1, 2}, {3}, {}, {}}, initialBoxes: []int{0}},
			16,
		},
		{
			params{status: []int{1, 0, 0, 0, 0, 0}, candies: []int{1, 1, 1, 1, 1, 1}, keys: [][]int{{1, 2, 3, 4, 5}, {}, {}, {}, {}, {}}, containedBoxes: [][]int{{1, 2, 3, 4, 5}, {}, {}, {}, {}, {}}, initialBoxes: []int{0}},
			6,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, maxCandies(p.status, p.candies, p.keys, p.containedBoxes, p.initialBoxes), ans)
	}
}
