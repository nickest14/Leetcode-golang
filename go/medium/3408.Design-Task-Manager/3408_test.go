package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem3408(t *testing.T) {
	obj := Constructor([][]int{{1, 101, 10}, {2, 102, 20}, {3, 103, 15}})
	obj.Add(4, 104, 5)
	obj.Edit(102, 8)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "ExecTop()", obj.ExecTop(), 3)
	obj.Rmv(101)
	obj.Add(5, 105, 15)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "ExecTop()", obj.ExecTop(), 5)
}
