// https://leetcode.com/problems/make-two-arrays-equal-by-reversing-subarrays/
// Level: Easy

package leetcode

func canBeEqual(target []int, arr []int) bool {
	if len(target) != len(arr) {
		return false
	}
	targetCount := make(map[int]int)
	arrCount := make(map[int]int)
	for _, num := range target {
		targetCount[num]++
	}
	for _, num := range arr {
		arrCount[num]++
	}
	for num, count := range targetCount {
		if count != arrCount[num] {
			return false
		}
	}

	return true
}

// // Another Solution
// func canBeEqual(target []int, arr []int) bool {
//     var counts [1001]int

//     for _, v := range target {
//         counts[v]++
//     }

//     for _, v := range arr {
//         counts[v]--
//         if counts[v] < 0 {
//             return false
//         }
//     }

//     return true
// }
