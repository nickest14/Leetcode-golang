// https://leetcode.com/problems/replace-non-coprime-numbers-in-array/
// Level: Hard

package leetcode

func replaceNonCoprimes(nums []int) []int {
	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	stack := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		stack = append(stack, nums[i])
		for len(stack) > 1 {
			num1 := stack[len(stack)-1]
			num2 := stack[len(stack)-2]
			g := gcd(num1, num2)
			if g > 1 {
				lcm := (num1 / g) * num2
				stack = stack[:len(stack)-2]
				stack = append(stack, lcm)
			} else {
				break
			}
		}
	}
	return stack
}
