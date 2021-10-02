package heapSort

import (
	"learnAlgorithm/sortTestHelper"
	"testing"
)

func TestSort(t *testing.T) {
	n := 1000000
	//n := 100
	arr := sortTestHelper.GenerateRandomArry(n, 0, n)
	arr2 := make([]int, len(arr))
	arr3 := make([]int, len(arr))
	copy(arr2, arr)
	copy(arr3, arr)
	//fmt.Println("before:", arr)
	//heapSort.HeapSort(arr)
	//fmt.Println("after:", arr)
	sortTestHelper.TestSort("heap sort1", Sort1, arr)
	sortTestHelper.TestSort("heap sort2", Sort2, arr2)
	sortTestHelper.TestSort("heap sort", Sort, arr3)
}
