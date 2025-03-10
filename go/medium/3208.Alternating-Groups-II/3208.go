// https://leetcode.com/problems/alternating-groups-ii/
// Level: Medium

package leetcode

func numberOfAlternatingGroups(colors []int, k int) int {
	colors = append(colors, colors[:k-1]...)
	left := 0
	ans := 0

	for right := 0; right < len(colors); right++ {
		if right > 0 && colors[right] == colors[right-1] {
			left = right
		}
		if right-left >= k-1 {
			ans++
		}
	}
	return ans
}
