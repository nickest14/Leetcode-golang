// https://leetcode.com/problems/construct-string-with-repeat-limit/
// Level: Medium

package leetcode

import (
	"container/heap"
	"strings"
)

type pair struct {
	ch    rune
	count int
}

type maxHeap []pair

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Less(i, j int) bool {
	return h[i].ch > h[j].ch
}

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(pair))
}

func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func repeatLimitedString(s string, repeatLimit int) string {
	freq := make(map[rune]int)
	for _, ch := range s {
		freq[ch]++
	}

	h := &maxHeap{}
	heap.Init(h)
	for ch, count := range freq {
		heap.Push(h, pair{ch: ch, count: count})
	}

	var ans strings.Builder
	for h.Len() > 0 {
		cur := heap.Pop(h).(pair)
		used := min(repeatLimit, cur.count)
		ans.WriteString(strings.Repeat(string(cur.ch), used))
		cur.count -= used

		if cur.count > 0 {
			if h.Len() == 0 {
				break
			}
			next := heap.Pop(h).(pair)
			ans.WriteString(string(next.ch))
			next.count--
			if next.count > 0 {
				heap.Push(h, next)
			}
			heap.Push(h, cur)
		}
	}

	return ans.String()
}
