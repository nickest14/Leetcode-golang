// https://leetcode.com/problems/find-smallest-letter-greater-than-target/
// Level: Easy

package leetcode

func nextGreatestLetter(letters []byte, target byte) byte {
	left, right := 0, len(letters)
	for left < right {
		mid := left + (right-left)/2
		if letters[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if left == len(letters) {
		return letters[0]
	}
	return letters[left]
}
