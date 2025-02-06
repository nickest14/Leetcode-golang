// https://leetcode.com/problems/tuple-with-same-product/
// Level: Medium

package leetcode

func tupleSameProduct(nums []int) int {
	count := make(map[int]int)
	ans := 0

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			prod := nums[i] * nums[j]
			ans += count[prod] * 8
			count[prod]++
		}
	}
	return ans
}
