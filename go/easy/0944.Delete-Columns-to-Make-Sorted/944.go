// https://leetcode.com/problems/delete-columns-to-make-sorted/
// Level: Easy

package leetcode

func minDeletionSize(strs []string) int {
	ans := 0
	for j := 0; j < len(strs[0]); j++ {
		for i := 1; i < len(strs); i++ {
			if strs[i][j] < strs[i-1][j] {
				ans++
				break
			}
		}
	}
	return ans
}
