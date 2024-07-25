// https://leetcode.com/problems/sort-an-array/
// Level: Medium

package leetcode

func sortArray(nums []int) []int {
	var heapify func(nums []int, n int, i int)
	heapify = func(nums []int, n int, i int) {
		largest := i
		left := 2*i + 1
		right := 2*i + 2
		if left < n && nums[largest] < nums[left] {
			largest = left
		}
		if right < n && nums[largest] < nums[right] {
			largest = right
		}
		if largest != i {
			nums[i], nums[largest] = nums[largest], nums[i]
			// heapify the root
			heapify(nums, n, largest)
		}
	}

	n := len(nums)
	// Build a maxheap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(nums, n, i)
	}
	// One by one extract elements
	for i := n - 1; i >= 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		heapify(nums, i, 0)
	}

	return nums
}
