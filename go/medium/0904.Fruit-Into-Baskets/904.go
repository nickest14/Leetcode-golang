// https://leetcode.com/problems/fruit-into-baskets/
// Level: Medium

package leetcode

func totalFruit(fruits []int) int {
	start := 0
	ans := 0
	fruitCount := make(map[int]int)

	for end := 0; end < len(fruits); end++ {
		fruitCount[fruits[end]]++

		for len(fruitCount) > 2 {
			fruitCount[fruits[start]]--
			if fruitCount[fruits[start]] == 0 {
				delete(fruitCount, fruits[start])
			}
			start++
		}

		ans = max(ans, end-start+1)
	}

	return ans
}
