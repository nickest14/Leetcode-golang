// https://leetcode.com/problems/check-if-array-pairs-are-divisible-by-k/
// Level: Medium

package leetcode

func canArrange(arr []int, k int) bool {
	if len(arr)%2 != 0 {
		return false
	}

	pairs := 0
	mapping := make(map[int]int)

	for _, val := range arr {
		key := (k - (val % k) + k) % k
		if mapping[key] > 0 {
			mapping[key]--
			pairs++
		} else {
			remainder := (val%k + k) % k
			mapping[remainder]++
		}
	}

	return pairs == len(arr)/2
}
