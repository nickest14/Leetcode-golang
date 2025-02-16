// https://leetcode.com/problems/find-the-punishment-number-of-an-integer/
// Level: Medium

package leetcode

func punishmentNumber(n int) int {
	var canPartition func(num, target int) bool
	canPartition = func(num, target int) bool {
		if target < 0 || num < target {
			return false
		}

		if num == target {
			return true
		}

		return canPartition(num/10, target-(num%10)) ||
			canPartition(num/100, target-(num%100)) ||
			canPartition(num/1000, target-(num%1000))
	}

	ans := 0

	for num := 1; num <= n; num++ {
		squr := num * num
		if canPartition(squr, num) {
			ans += squr
		}
	}
	return ans
}

// func punishmentNumber(n int) int {
// 	ans := 0
// 	memo := make(map[string]bool)

// 	var canPartition func(num int, target int) bool
// 	canPartition = func(num int, target int) bool {
// 		key := fmt.Sprintf("%d,%d", num, target)
// 		if val, exists := memo[key]; exists {
// 			return val
// 		}
// 		if target < 0 || num < target {
// 			return false
// 		}
// 		if num == target {
// 			return true
// 		}

// 		numStr := strconv.Itoa(num)
// 		for i := 0; i < len(numStr); i++ {
// 			left, _ := strconv.Atoi(numStr[:i+1])
// 			if target-left < 0 {
// 				continue
// 			}
// 			right, _ := strconv.Atoi(numStr[i+1:])
// 			if canPartition(right, target-left) {
// 				memo[key] = true
// 				return true
// 			}
// 		}

// 		memo[key] = false
// 		return false
// 	}

// 	for num := 1; num < n+1; num++ {
// 		squr := num * num
// 		if canPartition(squr, num) {
// 			ans += squr
// 		}
// 	}
// 	return ans
// }
