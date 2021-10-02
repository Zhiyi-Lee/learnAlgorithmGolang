package selectionSort

func Sort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		minIndex := i
		for j := i; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		if i != minIndex {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
}
