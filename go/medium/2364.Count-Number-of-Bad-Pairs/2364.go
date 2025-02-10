// https://leetcode.com/problems/count-number-of-bad-pairs/
// Level: Medium

package leetcode

func countBadPairs(nums []int) int64 {
	diffCount := make(map[int]int)
	badPairs := 0

	for i, num := range nums {
		val := num - i
		goodPairs := diffCount[val]
		badPairs += i - goodPairs
		diffCount[val]++
	}

	return int64(badPairs)
}
