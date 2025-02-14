// https://leetcode.com/problems/product-of-the-last-k-numbers/
// Level: Medium

package leetcode

type ProductOfNumbers struct {
	prod int
	li   []int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{prod: 1, li: []int{}}
}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.prod = 1
		this.li = []int{}
	} else {
		this.prod *= num
		this.li = append(this.li, this.prod)
	}
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	l := len(this.li)
	if l < k {
		return 0
	} else if l == k {
		return this.li[l-1]
	} else {
		return this.li[l-1] / this.li[l-1-k]
	}
}
