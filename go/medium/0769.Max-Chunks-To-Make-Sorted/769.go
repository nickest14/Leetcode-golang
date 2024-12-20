// https://leetcode.com/problems/max-chunks-to-make-sorted/
// Level: Medium

package leetcode

func maxChunksToSorted(arr []int) int {
	ans := 0
	maxSoFar := arr[0]
	for i, val := range arr {
		if maxSoFar < val {
			maxSoFar = val
		}
		if maxSoFar == i {
			ans++
		}
	}

	return ans
}
