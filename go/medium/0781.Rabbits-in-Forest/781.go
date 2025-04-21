// https://leetcode.com/problems/rabbits-in-forest/
// Level: Medium

package leetcode

func numRabbits(answers []int) int {
	ans := 0
	freq := make(map[int]int)
	for _, answer := range answers {
		freq[answer]++
	}
	for answer, count := range freq {
		groupSize := answer + 1
		groups := count / groupSize
		ans += groups * groupSize
		if count%groupSize > 0 {
			ans += groupSize
		}
	}
	return ans
}
