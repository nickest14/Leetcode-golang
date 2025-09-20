package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem3508(t *testing.T) {
	obj := Constructor(3)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "AddPacket()", obj.AddPacket(1, 4, 90), true)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "AddPacket()", obj.AddPacket(2, 5, 90), true)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "AddPacket()", obj.AddPacket(1, 4, 90), false)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "AddPacket()", obj.AddPacket(3, 5, 95), true)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "AddPacket()", obj.AddPacket(4, 5, 105), true)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "ForwardPacket()", obj.ForwardPacket(), []int{2, 5, 90})
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "AddPacket()", obj.AddPacket(5, 2, 110), true)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "GetCount()", obj.GetCount(5, 100, 110), 1)
}
