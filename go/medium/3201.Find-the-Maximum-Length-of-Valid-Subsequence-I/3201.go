// https://leetcode.com/problems/find-the-maximum-length-of-valid-subsequence-i/
// Level: Medium

package leetcode

func maximumLength(nums []int) int {
	countEven := 0
	countOdd := 0
	for _, num := range nums {
		if num%2 == 0 {
			countEven++
		} else {
			countOdd++
		}
	}

	evenDp := 0
	oddDp := 0

	for _, num := range nums {
		if num%2 == 0 {
			evenDp = max(evenDp, oddDp+1)
		} else {
			oddDp = max(oddDp, evenDp+1)
		}
	}

	return max(max(countEven, countOdd), max(evenDp, oddDp))
}
