// https://leetcode.com/problems/take-gifts-from-the-richest-pile/
// Level: Easy

package leetcode

import (
	"math"
	"sort"
)

func pickGifts(gifts []int, k int) int64 {
	sort.Ints(gifts)
	n := len(gifts)

	for i := 0; i < k; i++ {
		gifts[n-1] = int(math.Sqrt(float64(gifts[n-1])))
		sort.Ints(gifts)
	}

	ans := 0
	for _, gift := range gifts {
		ans += gift
	}
	return int64(ans)
}
