// https://leetcode.com/problems/find-lucky-integer-in-an-array/
// Level: Easy

package leetcode

func findLucky(arr []int) int {
	ans := -1
	count := make(map[int]int)
	for _, num := range arr {
		count[num]++
	}

	for num, freq := range count {
		if num == freq {
			ans = max(ans, num)
		}
	}
	return ans
}
