package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem2043(t *testing.T) {
	obj := Constructor([]int64{10, 100, 20, 50, 30})
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "withdraw(3, 10)", obj.Withdraw(3, 10), true)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "transfer(5, 1, 20)", obj.Transfer(5, 1, 20), true)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "deposit(5, 20)", obj.Deposit(5, 20), true)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "transfer(3, 4, 15)", obj.Transfer(3, 4, 15), false)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "withdraw(10, 50)", obj.Withdraw(10, 50), false)

}
