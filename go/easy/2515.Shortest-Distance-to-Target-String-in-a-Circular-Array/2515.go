// https://leetcode.com/problems/shortest-distance-to-target-string-in-a-circular-array/
// Level: Easy

package leetcode

func closestTarget(words []string, target string, startIndex int) int {
	n := len(words)
	for i := 0; i <= n/2; i++ {
		if words[(startIndex+i)%n] == target {
			return i
		}
		left := (startIndex - i) % n
		if left < 0 {
			left += n
		}
		if words[left] == target {
			return i
		}
	}
	return -1
}
