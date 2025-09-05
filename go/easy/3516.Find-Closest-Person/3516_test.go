package leetcode

import (
	"fmt"
	"testing"
)

type params struct {
	x int
	y int
	z int
}

type question3516 struct {
	params
	ans int
}

func Test_Problem3516(t *testing.T) {
	qs := []question3516{
		{
			params{x: 2, y: 7, z: 4},
			1,
		},
		{
			params{x: 2, y: 5, z: 6},
			2,
		},
		{
			params{x: 1, y: 5, z: 3},
			0,
		},
	}

	for _, q := range qs {
		p, ans := q.params, q.ans
		fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", p, findClosest(p.x, p.y, p.z), ans)
	}
}
