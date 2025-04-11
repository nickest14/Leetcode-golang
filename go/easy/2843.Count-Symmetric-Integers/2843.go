// https://leetcode.com/problems/count-symmetric-integers/
// Level: Easy

package leetcode

import "strconv"

func countSymmetricIntegers(low int, high int) int {
	ans := 0
	for num := low; num <= high; num++ {
		strNum := strconv.Itoa(num)
		if len(strNum)%2 == 0 {
			mid := len(strNum) / 2
			leftSum := 0
			rightSum := 0

			for i := 0; i < mid; i++ {
				leftSum += int(strNum[i] - '0')
			}
			for i := mid; i < len(strNum); i++ {
				rightSum += int(strNum[i] - '0')
			}

			if leftSum == rightSum {
				ans++
			}
		}
	}
	return ans
}
