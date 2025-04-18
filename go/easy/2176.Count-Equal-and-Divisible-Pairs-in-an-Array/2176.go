// https://leetcode.com/problems/count-equal-and-divisible-pairs-in-an-array/
// Level: Easy

package leetcode

func countPairs(nums []int, k int) int {
	ans := 0
	numToIndices := make(map[int][]int)
	for curIndex, num := range nums {
		if indices, ok := numToIndices[num]; ok {
			for _, index := range indices {
				if (curIndex*index)%k == 0 {
					ans++
				}
			}
		}
		numToIndices[num] = append(numToIndices[num], curIndex)
	}
	return ans
}
