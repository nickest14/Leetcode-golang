// https://leetcode.com/problems/water-bottles-ii/
// Level: Medium

package leetcode

func maxBottlesDrunk(numBottles int, numExchange int) int {
	ans := numBottles
	emptyBottles := numBottles

	for emptyBottles >= numExchange {
		emptyBottles -= numExchange
		numExchange++
		ans++
		emptyBottles++
	}
	return ans
}
