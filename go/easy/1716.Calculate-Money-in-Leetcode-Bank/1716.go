// https://leetcode.com/problems/calculate-money-in-leetcode-bank/
// Level: Easy

package leetcode

func totalMoney(n int) int {
	weeks := n / 7
	remain := n % 7
	fullWeeksSum := weeks*28 + 7*(weeks*(weeks-1)/2)
	remSum := remain*(1+weeks) + (remain * (remain - 1) / 2)
	return fullWeeksSum + remSum
}
