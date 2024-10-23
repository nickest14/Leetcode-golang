// https://leetcode.com/problems/kth-largest-sum-in-a-binary-tree/
// Level: Medium

package leetcode

import (
	"container/heap"
	structures "leetcode/go/structures"
)

// TreeNode is to define the struct
type TreeNode = structures.TreeNode

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

func heapReplace(h *MinHeap, val int) {
	(*h)[0] = val
	heap.Fix(h, 0)
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	q := []*TreeNode{root}
	minHeap := &MinHeap{}

	for len(q) > 0 {
		levelSum := 0
		levelSize := len(q)
		for i := 0; i < levelSize; i++ {
			node := q[0]
			q = q[1:]
			levelSum += node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		if minHeap.Len() < k {
			heap.Push(minHeap, levelSum)
		} else if levelSum > (*minHeap)[0] {
			heapReplace(minHeap, levelSum)
		}
	}

	if minHeap.Len() == k {
		return int64((*minHeap)[0])
	} else {
		return -1
	}
}
