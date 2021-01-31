package algorithms

func bubbleSort(arr []int) []int {
	arrLen := len(arr)

	for i := 0; i < arrLen; i++ {
		for j := 0; j < arrLen-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
