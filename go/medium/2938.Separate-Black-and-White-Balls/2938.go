// https://leetcode.com/problems/separate-black-and-white-balls/
// Level: Medium

package leetcode

func minimumSteps(s string) int64 {
	swap := 0
	black := 0
	for _, c := range s {
		if c == '0' {
			swap += black
		} else {
			black++
		}
	}
	return int64(swap)
}
