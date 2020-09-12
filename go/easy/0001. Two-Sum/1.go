// https://leetcode.com/problems/two-sum/
// Level: Easy

package leetcode

// TwoSum is to solve leetcode problem 1
func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for idx, val := range nums {
		if index, ok := m[target-val]; ok {
			return []int{idx, index}
		}
		m[val] = idx
	}
	return nil
}
