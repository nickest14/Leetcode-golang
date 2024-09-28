package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem641(t *testing.T) {
	obj := Constructor(3)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "InsertLast(1)", obj.InsertLast(1), true)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "InsertLast(2)", obj.InsertLast(2), true)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "InsertFront(3)", obj.InsertLast(3), true)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "InsertFront(4)", obj.InsertLast(4), false)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "IsEmpty()", obj.IsEmpty(), false)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "DeleteLast()", obj.DeleteLast(), true)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "InsertFront(4)", obj.InsertFront(4), true)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "GetFront()", obj.GetFront(), 4)
}
