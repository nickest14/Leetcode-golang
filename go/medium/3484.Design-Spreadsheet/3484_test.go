package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem3484(t *testing.T) {
	obj := Constructor(3)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "ExecTop()", obj.GetValue("=5+7"), 12)
	obj.SetCell("A1", 10)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "ExecTop()", obj.GetValue("=A1+6"), 16)
	obj.SetCell("B2", 15)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "ExecTop()", obj.GetValue("=A1+B2"), 25)
	obj.ResetCell("A1")
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "ExecTop()", obj.GetValue("=A1+B2"), 15)
}
