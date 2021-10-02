package heapSort

import (
	"fmt"
	maxHeap2 "learnAlgorithm/maxHeap"
)

func Sort1(arr []int) {
	maxHeap := maxHeap2.NewMaxHeap()

	for i := range arr {
		maxHeap.Insert(arr[i])
	}

	fmt.Println(maxHeap.Size())

	n := len(arr)
	for i := n - 1; i >= 0; i-- {
		item, has := maxHeap.ExtractMax()
		if has {
			arr[i] = item
			//fmt.Print(item, " ")
		}
	}
}

func Sort2(arr []int) {
	maxHeap := maxHeap2.NewMaxHeap()
	maxHeap.Heapify(arr)

	n := len(arr)
	for i := n - 1; i >= 0; i-- {
		item, has := maxHeap.ExtractMax()
		if has {
			arr[i] = item
			//fmt.Println(item)
		}
	}

}

//原地堆排序 索引从0开始
// parenr(i) = (i-1)/2
// left child (i) = 2*i + 1
// right child (i) = 2*i + 2
// 最后一个非叶子节点的索引 (count-1)/2
func Sort(arr []int) {
	n := len(arr)
	//heapity
	for i := (n - 1) / 2; i >= 0; i-- {
		shiftDown(arr, n, i)
	}

	//fmt.Println("data:", arr)
	//count := n
	for i := n - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		shiftDown(arr, i, 0)
	}
}

//索引从0开始
func shiftDown(arr []int, n int, k int) {
	for 2*k+1 < n {
		j := 2*k + 1 //在本轮循环中 arr[k] 和 arr[j] 交换位置
		if j+1 < n && arr[j+1] > arr[j] {
			j = j + 1
		}

		if arr[k] >= arr[j] {
			break
		}

		arr[k], arr[j] = arr[j], arr[k]
		k = j
	}
}
