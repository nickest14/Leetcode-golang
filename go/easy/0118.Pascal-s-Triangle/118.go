// https://leetcode.com/problems/pascals-triangle/
// Level: Easy

package leetcode

func generate(numRows int) [][]int {
	ans := [][]int{{1}} // layer 1

	for layer := 2; layer <= numRows; layer++ { // layer 2 to n
		tmp := []int{1}

		for i := 1; i < layer-1; i++ {
			tmp = append(tmp, ans[len(ans)-1][i]+ans[len(ans)-1][i-1])
		}

		tmp = append(tmp, 1)
		ans = append(ans, tmp)
	}

	return ans
}
