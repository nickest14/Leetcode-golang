// https://leetcode.com/problems/most-beautiful-item-for-each-query/
// Level: Medium

package leetcode

import "sort"

func maximumBeauty(items [][]int, queries []int) []int {
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] < items[j][0]
	})

	priceBeautyPairs := [][]int{}
	maxBeauty := 0
	for _, item := range items {
		price, beauty := item[0], item[1]
		if beauty > maxBeauty {
			maxBeauty = beauty
		}
		if len(priceBeautyPairs) == 0 || priceBeautyPairs[len(priceBeautyPairs)-1][1] < maxBeauty {
			priceBeautyPairs = append(priceBeautyPairs, []int{price, maxBeauty})
		}
	}
	ans := make([]int, len(queries))
	for i, q := range queries {
		idx := sort.Search(len(priceBeautyPairs), func(j int) bool {
			return priceBeautyPairs[j][0] > q
		}) - 1
		if idx >= 0 {
			ans[i] = priceBeautyPairs[idx][1]
		} else {
			ans[i] = 0
		}
	}

	return ans
}
