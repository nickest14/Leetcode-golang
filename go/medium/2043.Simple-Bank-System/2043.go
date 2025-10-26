// https://leetcode.com/problems/simple-bank-system/
// Level: Medium

package leetcode

type Bank struct {
	bal []int64
	n   int
}

func Constructor(balance []int64) Bank {
	return Bank{
		bal: append([]int64{}, balance...),
		n:   len(balance),
	}
}

func (this *Bank) valid(acc int) bool {
	return acc > 0 && acc <= this.n
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if !this.valid(account1) || !this.valid(account2) || this.bal[account1-1] < money {
		return false
	}
	this.bal[account1-1] -= money
	this.bal[account2-1] += money
	return true
}

func (this *Bank) Deposit(account int, money int64) bool {
	if !this.valid(account) {
		return false
	}
	this.bal[account-1] += money
	return true
}

func (this *Bank) Withdraw(account int, money int64) bool {
	if !this.valid(account) || this.bal[account-1] < money {
		return false
	}
	this.bal[account-1] -= money
	return true
}
