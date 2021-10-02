package mergeSort

import (
	"learnAlgorithm/insertionSort"
	"learnAlgorithm/sortTestHelper"
)

//合并索引为 [l,mid] [mid+1,r] 的数据
func __merge(arr []int, l int, mid int, r int) {
	aux := make([]int, r-l+1)
	//aux := arr[l:r+1] //不可以这样写啦 要新开一个用copy的 re-slice指向同一个底层数组哦
	copy(aux, arr[l:r+1])

	//mid := (l+r)/2
	i := l
	j := mid + 1
	for k := l; k <= r; k++ {
		if i > mid {
			arr[k] = aux[j-l]
			j++
		} else if j > r {
			arr[k] = aux[i-l]
			i++
		} else if aux[i-l] < aux[j-l] {
			arr[k] = aux[i-l]
			i++
		} else {
			arr[k] = aux[j-l]
			j++
		}
	}
}

//对索引为 [l,r] 区间的数据进行排序
func __mergeSort(arr []int, l int, r int) {
	if l >= r {
		return
	}

	mid := (l + r) / 2
	__mergeSort(arr, l, mid)
	__mergeSort(arr, mid+1, r)

	//__merge(arr,l,r)

	//对 nearly ordered 优化：[l,mid] [mid+1,r] 是有序的 如果 arr[mid] <= arr[mid+1] 则说明整个 [l,r] 是有序的 无需进行排序
	if arr[mid] > arr[mid+1] {
		__merge(arr, l, mid, r)
	}

}

//对索引为 [l,r] 区间的数据进行排序
//当分组len较小时 改用插入排序
func __mergeSortPro(arr []int, l int, r int) {
	if r-l <= 15 {
		insertionSort.SortSection(arr, l, r)
		return
	}

	mid := (l + r) / 2
	__mergeSort(arr, l, mid)
	__mergeSort(arr, mid+1, r)

	//__merge(arr,l,r)

	//对 nearly ordered 优化：[l,mid] [mid+1,r] 是有序的 如果 arr[mid] <= arr[mid+1] 则说明整个 [l,r] 是有序的 无需进行排序
	if arr[mid] > arr[mid+1] {
		__merge(arr, l, mid, r)
	}

}

func Sort(arr []int) {
	__mergeSort(arr, 0, len(arr)-1)
}

func SortPro(arr []int) {
	__mergeSortPro(arr, 0, len(arr)-1)
}

//自底向上的归并排序
func SortBU(arr []int) {
	n := len(arr)
	for size := 1; size <= n; size += size {
		for i := 0; i+size < n; i += 2 * size {
			__merge(arr, i, i+size-1, sortTestHelper.Min(i+2*size-1, n-1))
		}
	}
}

//自底向上的归并排序 Pro
func SortBUPro(arr []int) {
	n := len(arr)
	for size := 1; size <= n; size += size {
		//排到基本有序时 使用插入排序
		if n-size <= 16 {
			insertionSort.SortSection(arr, 0, n-1)
		} else {
			for i := 0; i+size < n; i += 2 * size {
				if arr[i+size-1] > arr[i+size] {
					__merge(arr, i, i+size-1, sortTestHelper.Min(i+2*size-1, n-1))
				}
			}
		}
	}
}
