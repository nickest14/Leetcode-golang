// https://leetcode.com/problems/shortest-subarray-to-be-removed-to-make-array-sorted/
// Level: Medium

package leetcode

func findLengthOfShortestSubarray(arr []int) int {
	n := len(arr)

	left := 0
	for left+1 < n && arr[left] <= arr[left+1] {
		left++
	}
	if left == n-1 {
		return 0
	}

	right := n - 1
	for right > 0 && arr[right] >= arr[right-1] {
		right--
	}

	ans := min(n-left-1, right)
	i, j := 0, right
	for i <= left && j < n {
		if arr[i] <= arr[j] {
			ans = min(ans, j-i-1)
			i++
		} else {
			j++
		}
	}

	return ans
}
