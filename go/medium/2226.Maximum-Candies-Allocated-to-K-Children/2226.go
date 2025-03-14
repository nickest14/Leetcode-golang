//  https://leetcode.com/problems/maximum-candies-allocated-to-k-children/
// Level: Medium

package leetcode

func maximumCandies(candies []int, k int64) int {
	ans := 0
	left, right := 0, 0

	for _, c := range candies {
		if c > right {
			right = c
		}
	}

	canDistribute := func(val int) bool {
		if val == 0 {
			return true
		}

		var count int64 = 0
		for _, c := range candies {
			count += int64(c / val)
		}
		return count >= k
	}

	for left <= right {
		mid := left + (right-left)/2
		if canDistribute(mid) {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return ans
}
