// https://leetcode.com/problems/find-elements-in-a-contaminated-binary-tree/
// Level: Medium

package leetcode

type FindSumPairs struct {
	map1 map[int]int
	map2 map[int]int
	arr2 []int
}

func Constructor(nums1 []int, nums2 []int) FindSumPairs {
	map1 := make(map[int]int)
	for _, num := range nums1 {
		map1[num]++
	}
	map2 := make(map[int]int)
	for _, num := range nums2 {
		map2[num]++
	}
	arr2 := make([]int, 0)
	for _, num := range nums2 {
		arr2 = append(arr2, num)
	}

	return FindSumPairs{
		map1: map1,
		map2: map2,
		arr2: arr2,
	}
}

func (this *FindSumPairs) Add(index int, val int) {
	oldVal := this.arr2[index]
	this.arr2[index] += val
	this.map2[oldVal]--
	this.map2[this.arr2[index]]++
}

func (this *FindSumPairs) Count(tot int) int {
	result := 0
	for num, count := range this.map1 {
		result += count * this.map2[tot-num]
	}
	return result
}
