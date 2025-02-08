// https://leetcode.com/problems/design-a-number-container-system/
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

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type NumberContainers struct {
	indexNum     map[int]int      // index -> number
	numToIndices map[int]*MinHeap // number -> minHeap of indices
}

func Constructor() NumberContainers {
	return NumberContainers{
		indexNum:     make(map[int]int),
		numToIndices: make(map[int]*MinHeap),
	}
}

func (this *NumberContainers) Change(index int, number int) {
	this.indexNum[index] = number

	if _, exists := this.numToIndices[number]; !exists {
		this.numToIndices[number] = &MinHeap{}
		heap.Init(this.numToIndices[number])
	}
	heap.Push(this.numToIndices[number], index)
}

func (this *NumberContainers) Find(number int) int {
	if h, exists := this.numToIndices[number]; exists {
		for h.Len() > 0 {
			minIndex := (*h)[0]
			if this.indexNum[minIndex] == number {
				return minIndex
			}
			heap.Pop(h)
		}
	}
	return -1
}
