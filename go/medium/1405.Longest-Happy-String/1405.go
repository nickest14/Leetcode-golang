// https://leetcode.com/problems/longest-happy-string/
// Level: Medium

package leetcode

import "container/heap"

type Item struct {
	freq int
	char byte
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].freq > pq[j].freq
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func longestDiverseString(a int, b int, c int) string {
	pq := &PriorityQueue{}
	heap.Init(pq)
	if a > 0 {
		heap.Push(pq, &Item{a, 'a'})
	}
	if b > 0 {
		heap.Push(pq, &Item{b, 'b'})
	}
	if c > 0 {
		heap.Push(pq, &Item{c, 'c'})
	}

	var ans []byte
	for pq.Len() > 1 {
		first := heap.Pop(pq).(*Item)
		if len(ans) < 2 || ans[len(ans)-1] != first.char || ans[len(ans)-2] != first.char {
			ans = append(ans, first.char)
			first.freq--
			if first.freq > 0 {
				heap.Push(pq, first)
			}
		} else {
			second := heap.Pop(pq).(*Item)
			ans = append(ans, second.char)
			second.freq--
			if second.freq > 0 {
				heap.Push(pq, second)
			}
			heap.Push(pq, first)
		}
	}
	if pq.Len() == 1 {
		last := heap.Pop(pq).(*Item)
		if last.freq <= 1 {
			ans = append(ans, last.char)
		} else {
			ans = append(ans, last.char, last.char)
		}
	}

	return string(ans)
}
