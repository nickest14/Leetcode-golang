// https://leetcode.com/problems/fruits-into-baskets-ii/
// Level: Easy

package leetcode

func numOfUnplacedFruits(fruits []int, baskets []int) int {
	basketUsed := make([]bool, len(baskets))
	unplacedCount := 0

	for _, fruit := range fruits {
		fruitPlaced := false
		for i := 0; i < len(baskets); i++ {
			if !basketUsed[i] && baskets[i] >= fruit {
				basketUsed[i] = true
				fruitPlaced = true
				break
			}
		}
		if !fruitPlaced {
			unplacedCount++
		}
	}

	return unplacedCount
}
