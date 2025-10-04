// https://leetcode.com/problems/water-bottles/
// Level: Easy

package leetcode

func numWaterBottles(numBottles int, numExchange int) int {
	ans := numBottles

	for numBottles >= numExchange {
		ans += numBottles / numExchange
		numBottles = numBottles/numExchange + numBottles%numExchange
	}
	return ans
}
