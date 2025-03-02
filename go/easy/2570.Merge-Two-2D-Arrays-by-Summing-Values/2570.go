// https://leetcode.com/problems/merge-two-2d-arrays-by-summing-values/
// Level: Easy

package leetcode

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	i, j := 0, 0
	ans := [][]int{}
	for i < len(nums1) && j < len(nums2) {
		if nums1[i][0] == nums2[j][0] {
			ans = append(ans, []int{nums1[i][0], nums1[i][1] + nums2[j][1]})
			i++
			j++
		} else if nums1[i][0] < nums2[j][0] {
			ans = append(ans, []int{nums1[i][0], nums1[i][1]})
			i++
		} else {
			ans = append(ans, []int{nums2[j][0], nums2[j][1]})
			j++
		}
	}
	for i < len(nums1) {
		ans = append(ans, nums1[i])
		i++
	}
	for j < len(nums2) {
		ans = append(ans, nums2[j])
		j++
	}

	return ans
}
