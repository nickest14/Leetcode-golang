// https://leetcode.com/problems/find-all-people-with-secret/
// Level: Hard

package leetcode

import "sort"

func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(a, b int) {
		parent[find(b)] = find(a)
	}

	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][2] < meetings[j][2]
	})

	union(0, firstPerson)

	i := 0
	for i < len(meetings) {
		time := meetings[i][2]
		involved := make(map[int]bool)

		for i < len(meetings) && meetings[i][2] == time {
			x, y := meetings[i][0], meetings[i][1]
			union(x, y)
			involved[x] = true
			involved[y] = true
			i++
		}

		root0 := find(0)
		for p := range involved {
			if find(p) != root0 {
				parent[p] = p
			}
		}
	}

	var ans []int
	for i := 0; i < n; i++ {
		if find(i) == find(0) {
			ans = append(ans, i)
		}
	}
	return ans
}
