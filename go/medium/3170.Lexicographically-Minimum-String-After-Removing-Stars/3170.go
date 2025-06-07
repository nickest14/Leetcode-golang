// https://leetcode.com/problems/lexicographically-minimum-string-after-removing-stars/
// Level: Medium

package leetcode

import "container/heap"

type CharIndex struct {
	char  rune
	index int
}

type MinHeap []CharIndex

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	if h[i].char == h[j].char {
		return h[i].index < h[j].index
	}
	return h[i].char < h[j].char
}
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(CharIndex))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func clearStars(s string) string {
	h := &MinHeap{}
	heap.Init(h)

	for i, c := range s {
		if c == '*' {
			if h.Len() > 0 {
				heap.Pop(h)
			}
		} else {
			heap.Push(h, CharIndex{c, -i})
		}
	}

	char := make([]rune, len(s))
	for h.Len() > 0 {
		item := heap.Pop(h).(CharIndex)
		char[-item.index] = item.char
	}

	ans := make([]rune, 0)
	for _, v := range char {
		if v != 0 {
			ans = append(ans, v)
		}
	}
	return string(ans)
}
