// https://leetcode.com/problems/sort-the-people/
// Level: Easy

package leetcode

import (
	"sort"
)

type Person struct {
	Name   string
	Height int
}

func sortPeople(names []string, heights []int) []string {
	length := len(names)
	people := make([]Person, length)

	for i := 0; i < length; i++ {
		people[i] = Person{Name: names[i], Height: heights[i]}
	}
	sort.Slice(people, func(i, j int) bool {
		return people[i].Height > people[j].Height
	})

	ans := make([]string, length)
	for i := 0; i < length; i++ {
		ans[i] = people[i].Name
	}

	return ans
}
