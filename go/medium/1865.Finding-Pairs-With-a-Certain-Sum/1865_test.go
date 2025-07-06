package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem1865(t *testing.T) {
	obj := Constructor([]int{1, 1, 2, 2, 2, 3}, []int{1, 4, 5, 2, 5, 4})
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "Count(7)", obj.Count(7), 8)
	obj.Add(3, 2)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "Count(8)", obj.Count(8), 2)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "Count(4)", obj.Count(4), 1)
	obj.Add(0, 1)
	obj.Add(1, 1)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "Count(7)", obj.Count(7), 11)
	fmt.Println()
}
