// https://leetcode.com/problems/container-with-most-water/
// Level: Medium

package leetcode

func maxArea(height []int) int {
	ans := 0
	l, r := 0, len(height)-1

	h := height[0]
	for _, num := range height[1:] {
		if num > h {
			h = num
		}
	}

	for l < r {
		ans = max(ans, min(height[l], height[r])*(r-l))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
		if (r-l)*h <= ans {
			break
		}
	}
	return ans
}
