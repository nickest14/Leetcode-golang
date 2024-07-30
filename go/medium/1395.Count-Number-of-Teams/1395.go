// https://leetcode.com/problems/count-number-of-teams/
// Level: Medium

package leetcode

func numTeams(rating []int) int {
	var ans int
	n := len(rating)
	for mid := range n {
		var leftSmaller, rightLarger int = 0, 0
		for left := 0; left < mid; left++ {
			if rating[left] < rating[mid] {
				leftSmaller++
			}
		}
		for right := mid + 1; right < n; right++ {
			if rating[right] > rating[mid] {
				rightLarger++
			}
		}

		// Calculate and add the number of ascending rating teams (small-mid-large)
		ans += leftSmaller * rightLarger

		leftLarger := mid - leftSmaller
		rightSmaller := n - 1 - mid - rightLarger
		// Calculate and add the number of descending rating teams (large-mid-small)
		ans += leftLarger * rightSmaller
	}
	return ans
}
