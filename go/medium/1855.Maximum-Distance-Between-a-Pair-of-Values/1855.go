// https://leetcode.com/problems/maximum-distance-between-a-pair-of-values/
// Level: Medium

package leetcode

func maxDistance(nums1 []int, nums2 []int) int {
	i, j := 0, 1
	n1, n2 := len(nums1), len(nums2)
	for i < n1 && j < n2 {
		if nums1[i] > nums2[j] {
			i++
		}
		j++
	}
	return j - i - 1
}
