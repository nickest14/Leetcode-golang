// https://leetcode.com/problems/bitwise-ors-of-subarrays/
// Level: Medium

package leetcode

func subarrayBitwiseORs(arr []int) int {
	ans := make(map[int]bool)
	cur := make(map[int]bool)

	for _, num := range arr {
		next := make(map[int]bool)
		next[num] = true
		for x := range cur {
			next[x|num] = true
		}
		cur = next
		for x := range cur {
			ans[x] = true
		}
	}

	return len(ans)
}
