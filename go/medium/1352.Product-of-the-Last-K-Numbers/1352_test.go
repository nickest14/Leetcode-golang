package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem1352(t *testing.T) {
	obj := Constructor()
	obj.Add(3)
	obj.Add(0)
	obj.Add(2)
	obj.Add(5)
	obj.Add(4)

	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "Find()", obj.GetProduct(2), 20)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "Find()", obj.GetProduct(3), 40)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "Find()", obj.GetProduct(4), 0)
}
