// https://leetcode.com/problems/count-number-of-trapezoids-i/
// Level: Medium

package leetcode

func countTrapezoids(points [][]int) int {
	const mod = 1_000_000_007
	mp := map[int]int{}
	for _, p := range points {
		mp[p[1]]++
	}

	seg := []int{}
	for _, v := range mp {
		if v >= 2 {
			seg = append(seg, v*(v-1)/2%mod)
		}
	}
	sum, ans := 0, 0
	for _, v := range seg {
		ans = (ans + v*sum) % mod
		sum = (sum + v) % mod
	}
	return ans
}
