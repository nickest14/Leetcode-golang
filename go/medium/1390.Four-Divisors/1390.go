// https://leetcode.com/problems/four-divisors/
// Level: Medium

package leetcode

func sumFourDivisors(nums []int) int {
	ans := 0

	for _, num := range nums {
		divCount := 0
		inSum := 0

		for d := 1; d*d <= num; d++ {
			if num%d == 0 {
				other := num / d

				if d == other {
					divCount++
					inSum += d
				} else {
					divCount += 2
					inSum += d + other
				}

				if divCount > 4 {
					break
				}
			}
		}

		if divCount == 4 {
			ans += inSum
		}
	}

	return ans
}
