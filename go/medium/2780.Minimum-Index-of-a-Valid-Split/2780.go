// https://leetcode.com/problems/minimum-index-of-a-valid-split/
// Level: Medium

package leetcode

func minimumIndex(nums []int) int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	dominant := 0
	count := 0
	for num, c := range freq {
		if c > len(nums)/2 {
			dominant = num
			count = c
			break
		}
	}

	leftCount := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == dominant {
			leftCount++
		}
		leftSize := i + 1
		rightSize := len(nums) - leftSize
		rightCount := count - leftCount

		if leftCount*2 > leftSize && rightCount*2 > rightSize {
			return i
		}
	}

	return -1
}
