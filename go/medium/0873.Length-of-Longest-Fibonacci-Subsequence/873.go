// https://leetcode.com/problems/length-of-longest-fibonacci-subsequence/
// Level: Medium

package leetcode

func lenLongestFibSubseq(arr []int) int {
	lookup := make(map[int]map[int]int)
	maxLength := 0

	for curIndex, curVal := range arr {
		lookup[curVal] = make(map[int]int)
		for prevIndex := curIndex - 1; prevIndex >= 0; prevIndex-- {
			prevVal := arr[prevIndex]
			prev2Val := curVal - prevVal
			if prev2Val >= prevVal {
				break
			}

			if _, exists := lookup[prev2Val]; exists {
				if length, found := lookup[prevVal][prev2Val]; found {
					lookup[curVal][prevVal] = length + 1
				} else {
					lookup[curVal][prevVal] = 2 + 1
				}
				if lookup[curVal][prevVal] > maxLength {
					maxLength = lookup[curVal][prevVal]
				}
			}
		}
	}

	return maxLength
}
