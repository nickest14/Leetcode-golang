package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem1381(t *testing.T) {
	obj := Constructor(3)
	obj.Push(1)
	obj.Push(2)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "Pop()", obj.Pop(), 2)
	obj.Push(2)
	obj.Push(3)
	obj.Push(4)
	obj.Increment(5, 100)
	obj.Increment(2, 100)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "Pop()", obj.Pop(), 103)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "Pop()", obj.Pop(), 202)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "Pop()", obj.Pop(), 201)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "Pop()", obj.Pop(), -1)
}
