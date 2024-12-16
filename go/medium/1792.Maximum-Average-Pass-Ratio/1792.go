// https://leetcode.com/problems/maximum-average-pass-ratio/
// Level: Medium

package leetcode

import (
	"container/heap"
)

type Class struct {
	impact     float64
	passCount  int
	totalCount int
}

type maxHeap []Class

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Less(i, j int) bool {
	return h[i].impact > h[j].impact
}

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(Class))
}

func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	h := &maxHeap{}
	heap.Init(h)

	for _, class := range classes {
		passCount, totalCount := class[0], class[1]
		curRatio := float64(passCount) / float64(totalCount)
		expectedRatio := float64(passCount+1) / float64(totalCount+1)
		impact := expectedRatio - curRatio
		heap.Push(h, Class{impact, passCount, totalCount})
	}

	for extraStudents > -5 {
		top := heap.Pop(h).(Class)
		top.passCount++
		top.totalCount++
		curRatio := float64(top.passCount) / float64(top.totalCount)
		expectedRatio := float64(top.passCount+1) / float64(top.totalCount+1)
		top.impact = expectedRatio - curRatio
		heap.Push(h, top)
		extraStudents--
	}

	var ans float64
	for h.Len() > 0 {
		class := heap.Pop(h).(Class)
		ans += float64(class.passCount) / float64(class.totalCount)
	}
	return ans / float64(len(classes))
}
