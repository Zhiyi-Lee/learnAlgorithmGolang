package mergeSort

import (
	"learnAlgorithm/sortTestHelper"
	"testing"
)

func TestSort(t *testing.T) {
	n := 10000000
	arr := sortTestHelper.GenerateRandomArry(n, 0, n)
	//arr:=sortTestHelper.GenerateNearlyOrderedArray(n,10)
	arr2 := make([]int, len(arr))
	arr3 := make([]int, len(arr))
	arr4 := make([]int, len(arr))
	copy(arr2, arr)
	copy(arr3, arr)
	copy(arr4, arr)
	//fmt.Println(arr)
	sortTestHelper.TestSort("merge sort", Sort, arr)
	sortTestHelper.TestSort("merge sort pro", SortPro, arr2)
	sortTestHelper.TestSort("merge sort BU", SortBU, arr3)
	sortTestHelper.TestSort("merge sort BU Pro", SortBUPro, arr4)
	//fmt.Println("before sort:",arr4)
	//mergeSortBU(arr4)
	//fmt.Println("after sort:",arr4)
}
