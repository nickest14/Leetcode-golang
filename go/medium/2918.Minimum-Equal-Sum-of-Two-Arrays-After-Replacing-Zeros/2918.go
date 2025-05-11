// https://leetcode.com/problems/minimum-equal-sum-of-two-arrays-after-replacing-zeros/
// Level: Medium

package leetcode

func minSum(nums1 []int, nums2 []int) int64 {
	var nums1Sum, nums2Sum int
	var nums1Zeros, nums2Zeros int
	for _, num := range nums1 {
		if num == 0 {
			nums1Zeros++
			nums1Sum++
		}
		nums1Sum += num
	}
	for _, num := range nums2 {
		if num == 0 {
			nums2Zeros++
			nums2Sum++
		}
		nums2Sum += num
	}

	if (nums1Zeros == 0 && nums2Sum > nums1Sum) || (nums2Zeros == 0 && nums1Sum > nums2Sum) {
		return -1
	}

	if nums1Sum > nums2Sum {
		return int64(nums1Sum)
	}
	return int64(nums2Sum)
}
