// https://leetcode.com/problems/design-a-food-rating-system/
// Level: Medium

package leetcode

import (
	"container/heap"
)

type FoodItem struct {
	name   string
	rating int
}

type PriorityQueue []FoodItem

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].rating == pq[j].rating {
		return pq[i].name < pq[j].name
	}
	return pq[i].rating > pq[j].rating
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(FoodItem))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

type FoodRatings struct {
	foodToCuisine map[string]string
	foodToRating  map[string]int
	cuisineHeaps  map[string]*PriorityQueue
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	fr := FoodRatings{
		foodToCuisine: make(map[string]string),
		foodToRating:  make(map[string]int),
		cuisineHeaps:  make(map[string]*PriorityQueue),
	}

	for i, food := range foods {
		fr.foodToCuisine[food] = cuisines[i]
		fr.foodToRating[food] = ratings[i]

		if fr.cuisineHeaps[cuisines[i]] == nil {
			pq := &PriorityQueue{}
			heap.Init(pq)
			fr.cuisineHeaps[cuisines[i]] = pq
		}
		heap.Push(fr.cuisineHeaps[cuisines[i]], FoodItem{food, ratings[i]})
	}
	return fr
}

func (this *FoodRatings) ChangeRating(food string, newRating int) {
	cuisine := this.foodToCuisine[food]
	this.foodToRating[food] = newRating
	heap.Push(this.cuisineHeaps[cuisine], FoodItem{food, newRating})
}

func (this *FoodRatings) HighestRated(cuisine string) string {
	pq := this.cuisineHeaps[cuisine]

	for pq.Len() > 0 {
		top := (*pq)[0]
		if this.foodToRating[top.name] == top.rating {
			return top.name
		}
		heap.Pop(pq)
	}
	return ""
}
