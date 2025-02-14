// https://leetcode.com/problems/minimum-operations-to-exceed-threshold-value-ii/
// Level: Medium

package leetcode

import "container/heap"

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func minOperations(nums []int, k int) int {
	h := MinHeap(nums)
	heap.Init(&h)
	ans := 0

	for len(h) > 1 {
		x := heap.Pop(&h).(int)
		y := heap.Pop(&h).(int)
		if x >= k {
			return ans
		}
		heap.Push(&h, x*2+y)
		ans++
	}

	if heap.Pop(&h).(int) >= k {
		return ans
	}

	return -1
}
