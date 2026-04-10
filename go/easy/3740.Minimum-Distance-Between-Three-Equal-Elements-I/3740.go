// https://leetcode.com/problems/minimum-distance-between-three-equal-elements-i/
// Level: Easy

package leetcode

func minimumDistance(nums []int) int {
	mp := make(map[int][]int, len(nums))
	for i, v := range nums {
		mp[v] = append(mp[v], i)
	}

	const inf = int(^uint(0) >> 1)
	ans := inf
	for _, indices := range mp {
		if len(indices) < 3 {
			continue
		}
		for i := 0; i+2 < len(indices); i++ {
			dist := 2 * (indices[i+2] - indices[i])
			if dist < ans {
				ans = dist
			}
		}
	}

	if ans == inf {
		return -1
	}
	return ans
}
