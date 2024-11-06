// https://leetcode.com/problems/find-if-array-can-be-sorted/
// Level: Medium

package leetcode

func canSortArray(nums []int) bool {
	var (
		prevMax, iSetBit = -1, -1
		iMin, iMax       = 0, -1
	)

	for _, num := range nums {
		curSetBit := getBitCount(num)
		if curSetBit == iSetBit {
			// current interval continues
			iMin, iMax = min(iMin, num), max(iMax, num)
		} else {
			// new interval begins
			if prevMax != -1 && prevMax > iMin {
				return false
			}
			prevMax = iMax
			iMin, iMax = num, num
			iSetBit = curSetBit
		}
	}

	return prevMax == -1 || prevMax < iMin
}

func getBitCount(num int) int {
	count := 0
	for num > 0 {
		count += num & 1
		num >>= 1
	}
	return count
}
