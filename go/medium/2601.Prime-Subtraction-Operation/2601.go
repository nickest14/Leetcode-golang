// https://leetcode.com/problems/prime-subtraction-operation/
// Level: Medium

package leetcode

import (
	"sort"
)

func primeSubOperation(nums []int) bool {
	processPrime := func(maxNum int) []int {
		isPrime := make([]bool, maxNum+1)
		for i := range isPrime {
			isPrime[i] = true
		}
		isPrime[0], isPrime[1] = false, false
		for i := 2; i*i <= maxNum; i++ {
			if isPrime[i] {
				for j := i * i; j <= maxNum; j += i {
					isPrime[j] = false
				}
			}
		}
		primes := []int{}
		for i := 2; i <= maxNum; i++ {
			if isPrime[i] {
				primes = append(primes, i)
			}
		}
		return primes
	}

	primes := processPrime(1000)
	prev := 0
	for i := 0; i < len(nums); i++ {
		pos := sort.Search(len(primes), func(j int) bool {
			return primes[j] >= nums[i]
		})
		success := false
		for j := pos - 1; j >= 0; j-- {
			prime := primes[j]
			if nums[i]-prime > prev {
				nums[i] -= prime
				success = true
				break
			}
		}
		if !success && nums[i] <= prev {
			return false
		}
		prev = nums[i]
	}

	return true
}
