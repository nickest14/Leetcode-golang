// https://leetcode.com/problems/xor-queries-of-a-subarray/
// Level: Medium

package leetcode

import "fmt"

func xorQueries(arr []int, queries [][]int) []int {
	length := len(arr)
	prefixs := make([]int, length)
	ans := make([]int, 0, len(queries))

	prefixs[0] = arr[0]
	for i := 1; i < length; i++ {
		prefixs[i] = prefixs[i-1] ^ arr[i]
	}

	for _, query := range queries {
		left, right := query[0], query[1]
		if left == 0 {
			ans = append(ans, prefixs[right])
		} else {
			fmt.Println(prefixs[right], prefixs[left-1])
			ans = append(ans, prefixs[right]^prefixs[left-1])
		}
	}
	return ans
}
