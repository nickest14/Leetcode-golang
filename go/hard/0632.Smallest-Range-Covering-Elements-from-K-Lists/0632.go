// https://leetcode.com/problems/smallest-range-covering-elements-from-k-lists/
// Level: Hard

package leetcode

import (
	"container/heap"
	"math"
)

type Element struct {
	val, row, col int
}

type MinHeap []Element

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].val < h[j].val
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Element))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func smallestRange(nums [][]int) []int {
	n := len(nums)
	h := &MinHeap{}
	ma := math.Inf(-1)

	heap.Init(h)
	for i := 0; i < n; i++ {
		heap.Push(h, Element{val: nums[i][0], row: i, col: 0})
		if float64(nums[i][0]) > ma {
			ma = float64(nums[i][0])
		}
	}

	ans := []int{(*h)[0].val, int(ma)}
	for {
		minElement := heap.Pop(h).(Element)
		minVal, i, j := minElement.val, minElement.row, minElement.col
		if int(ma)-minVal < ans[1]-ans[0] {
			ans[0], ans[1] = minVal, int(ma)
		}
		if j == len(nums[i])-1 {
			break
		}
		nextNum := nums[i][j+1]
		heap.Push(h, Element{val: nextNum, row: i, col: j + 1})
		if float64(nextNum) > ma {
			ma = float64(nextNum)
		}
	}

	return ans
}
