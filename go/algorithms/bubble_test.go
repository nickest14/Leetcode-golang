package algorithms

import (
	"fmt"
	"testing"
)

type BubbleSort struct {
	param []int
	ans   []int
}

func Test_BubbleSort(t *testing.T) {
	qs := []BubbleSort{
		{
			[]int{3, 2, 1},
			[]int{1, 2, 3},
		},
		{
			[]int{10, 1, 3, 9, 4, 5, 2, 6},
			[]int{1, 2, 3, 4, 5, 6, 9, 10},
		},
	}
	for _, q := range qs {
		p, ans := q.param, q.ans
		fmt.Printf("[input]: %v        ", p)
		fmt.Printf("[output]: %v    [answer]: %v\n", bubbleSort(p), ans)
	}
}
