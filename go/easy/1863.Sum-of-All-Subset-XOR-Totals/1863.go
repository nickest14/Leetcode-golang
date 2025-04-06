// https://leetcode.com/problems/sum-of-all-subset-xor-totals/
// Level: Easy

package leetcode

func subsetXORSum(nums []int) int {
	var dfs func(index int, xor int) int
	dfs = func(i int, xor int) int {
		if i == len(nums) {
			return xor
		}

		includeNum := dfs(i+1, xor^nums[i])
		excludeNum := dfs(i+1, xor)
		return includeNum + excludeNum
	}
	return dfs(0, 0)
}
