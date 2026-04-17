package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	nums []int
}

type question3761 struct {
	params
	ans int
}

func Test_Problem3761(t *testing.T) {
	qs := []question3761{
		{
			params{nums: []int{12, 21, 45, 33, 54}},
			1,
		},
		{
			params{nums: []int{120, 21}},
			1,
		},
		{
			params{nums: []int{3450, 1591, 8381, 5485, 646, 9444, 3832, 1263, 1529, 7577, 9421, 3080, 5084, 7797, 9527, 3845, 603, 4780, 6253, 5787, 1993, 6204, 7680, 9146, 1838, 5510, 2339, 5751}},
			22,
		},
	}
	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, minMirrorPairDistance(p.nums), ans)
	}
}
