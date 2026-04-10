// https://leetcode.com/problems/xor-after-range-multiplication-queries-i/
// Level: Medium

package leetcode

func xorAfterQueries(nums []int, queries [][]int) int {
	const mod int64 = 1_000_000_007
	for _, q := range queries {
		l, r, k, v := q[0], q[1], q[2], q[3]
		for idx := l; idx <= r; idx += k {
			nums[idx] = int((int64(nums[idx]) * int64(v)) % mod)
		}
	}
	ans := 0
	for _, num := range nums {
		ans ^= num
	}
	return ans
}
