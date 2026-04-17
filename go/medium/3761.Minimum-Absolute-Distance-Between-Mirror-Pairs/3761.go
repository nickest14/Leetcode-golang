// https://leetcode.com/problems/minimum-absolute-distance-between-mirror-pairs/
// Level: Medium

package leetcode

func minMirrorPairDistance(nums []int) int {
	reverseInt := func(x int) int {
		r := 0
		for x > 0 {
			r = r*10 + x%10
			x /= 10
		}
		return r
	}

	const inf = 100000
	ans := inf
	seen := make(map[int]int, len(nums))

	for i, num := range nums {
		if j, ok := seen[num]; ok {
			if d := i - j; d < ans {
				ans = d
			}
		}
		seen[reverseInt(num)] = i
	}

	if ans == inf {
		return -1
	}
	return ans
}
