// https://leetcode.com/problems/soup-servings/
// Level: Medium

package leetcode

import (
	"fmt"
	"math"
)

func soupServings(n int) float64 {
	if n >= 4800 {
		return 1.0
	}

	n = int(math.Ceil(float64(n) / 25.0))
	memo := make(map[string]float64)

	var dfs func(a, b int) float64
	dfs = func(a, b int) float64 {
		key := fmt.Sprintf("%d,%d", a, b)
		if val, exists := memo[key]; exists {
			return val
		}

		if a == 0 && b == 0 {
			return 0.5
		}
		if a == 0 {
			return 1.0
		}
		if b == 0 {
			return 0.0
		}

		// Four operations
		op1 := dfs(max(0, a-4), b)
		op2 := dfs(max(0, a-3), max(0, b-1))
		op3 := dfs(max(0, a-2), max(0, b-2))
		op4 := dfs(max(0, a-1), max(0, b-3))

		ans := (op1 + op2 + op3 + op4) * 0.25
		memo[key] = ans
		return ans
	}

	return dfs(n, n)
}
