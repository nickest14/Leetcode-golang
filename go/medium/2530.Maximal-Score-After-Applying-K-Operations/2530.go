// https://leetcode.com/problems/maximal-score-after-applying-k-operations/
// Level: Medium

package leetcode

import (
	"container/heap"
	"math"
)

type maxHeap []int

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func maxKelements(nums []int, k int) int64 {
	maxHeap := &maxHeap{}
	heap.Init(maxHeap)
	for _, num := range nums {
		heap.Push(maxHeap, num)
	}

	var ans int64 = 0
	for k > 0 {
		val := heap.Pop(maxHeap).(int)
		ans += int64(val)
		newVal := int(math.Ceil(float64(val) / 3))
		heap.Push(maxHeap, newVal)
		k--
	}
	return ans
}
