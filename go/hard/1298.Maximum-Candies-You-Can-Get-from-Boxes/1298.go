// https://leetcode.com/problems/maximum-candies-you-can-get-from-boxes/
// Level: Hard

package leetcode

func maxCandies(status []int, candies []int, keys [][]int, containedBoxes [][]int, initialBoxes []int) int {
	q := []int{}
	for _, box_id := range initialBoxes {
		for _, key := range keys[box_id] {
			status[key] = 1
		}

		if status[box_id] == 0 {
			q = append(q, box_id)
		} else {
			q = append([]int{box_id}, q...)
		}
	}

	ans := 0
	for len(q) > 0 {
		ind := q[0]
		q = q[1:]

		if status[ind] == 0 {
			break
		} else if status[ind] == 1 {
			ans += candies[ind]
			for _, key := range keys[ind] {
				status[key] = 1
			}
		}

		for _, new_box_ind := range containedBoxes[ind] {
			if status[new_box_ind] == 0 {
				q = append(q, new_box_ind)
			} else {
				q = append([]int{new_box_ind}, q...)
			}
		}
	}
	return ans
}
