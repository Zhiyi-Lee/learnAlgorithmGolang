package insertionSort

func Sort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}

}

func SortPro(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		e := arr[i]
		var j int
		for j = i; j > 0 && arr[j-1] > e; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = e
	}
}

//对slice [l,r] 进行直接插入排序
func SortSection(arr []int, l int, r int) {
	for i := l; i <= r; i++ {
		e := arr[i]
		var j int
		for j = i; j > l && e < arr[j-1]; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = e
	}
}

//希尔排序
func ShellSort(arr []int) {
	n := len(arr)
	for step := n / 2; step > 0; step /= 2 {
		for i := step; i < n; i++ {
			for j := i - step; j >= 0 && arr[j] > arr[j+step]; j -= step {
				arr[j], arr[j+step] = arr[j+step], arr[j]
			}
		}
	}
}

//希尔排序 pro
func ShellSortPro(arr []int) {
	n := len(arr)
	for step := n / 2; step > 0; step /= 2 {
		for i := step; i < n; i++ {
			e := arr[i]
			var j int
			for j = i - step; j >= 0 && arr[j] > e; j -= step {
				arr[j+step] = arr[j]
			}
			arr[j+step] = e
		}
	}
}
