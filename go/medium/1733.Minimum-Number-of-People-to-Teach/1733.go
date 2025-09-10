// https://leetcode.com/problems/minimum-number-of-people-to-teach/
// Level: Medium

package leetcode

func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
	usersToTeach := make(map[int]bool)

	for _, friendship := range friendships {
		u1, u2 := friendship[0]-1, friendship[1]-1
		canCommunicate := false

		for _, lang1 := range languages[u1] {
			for _, lang2 := range languages[u2] {
				if lang1 == lang2 {
					canCommunicate = true
					break
				}
			}
			if canCommunicate {
				break
			}
		}

		if !canCommunicate {
			usersToTeach[u1] = true
			usersToTeach[u2] = true
		}
	}

	ans := len(languages) + 1
	for language := 1; language <= n; language++ {
		count := 0
		for user := range usersToTeach {
			knows := false
			for _, lang := range languages[user] {
				if lang == language {
					knows = true
					break
				}
			}
			if !knows {
				count++
			}
		}
		if count < ans {
			ans = count
		}
	}

	return ans
}
