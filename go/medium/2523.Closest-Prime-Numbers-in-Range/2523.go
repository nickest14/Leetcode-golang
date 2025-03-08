// https://leetcode.com/problems/closest-prime-numbers-in-range/
// Level: Medium

package leetcode

func closestPrimes(left int, right int) []int {
	primes := make([]bool, right+1)
	for i := range primes {
		primes[i] = true
	}
	primes[0], primes[1] = false, false

	for i := 2; i*i <= right; i++ {
		if primes[i] {
			for j := i * i; j <= right; j += i {
				primes[j] = false
			}
		}
	}

	primesInRange := []int{}
	for i := left; i <= right; i++ {
		if primes[i] {
			primesInRange = append(primesInRange, i)
		}
	}

	ans := []int{-1, -1}
	if len(primesInRange) < 2 {
		return ans
	}

	minGap := int(^uint(0) >> 1)

	for i := 1; i < len(primesInRange); i++ {
		gap := primesInRange[i] - primesInRange[i-1]
		if gap < minGap {
			minGap = gap
			ans[0] = primesInRange[i-1]
			ans[1] = primesInRange[i]
		}
	}

	return ans
}
