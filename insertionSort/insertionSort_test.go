package insertionSort

import (
	"learnAlgorithm/sortTestHelper"
	"testing"
)

func TestSort(t *testing.T) {
	n := 100000
	arr := sortTestHelper.GenerateRandomArry(n, 0, n)
	arr2 := make([]int, len(arr))
	arr3 := make([]int, len(arr))
	arr4 := make([]int, len(arr))
	copy(arr2, arr)
	copy(arr3, arr)
	copy(arr4, arr)
	//fmt.Println("before sore:",arr4)
	//insertionSortPro( arr)
	//shellSort(arr4)
	//shellSortPro(arr4)
	//fmt.Println("after sort:",arr4)

	//fmt.Println("n=", n)
	sortTestHelper.TestSort("insertion sort", Sort, arr)
	sortTestHelper.TestSort("insertion sort pro", SortPro, arr2)
	sortTestHelper.TestSort("shell sort", ShellSort, arr3)
	sortTestHelper.TestSort("shell sort pro", ShellSortPro, arr4)
	//fmt.Println(arr2)

	//fmt.Println("sort a neearly ordered arry")
	//arr2:=sortTestHelper.GenerateNearlyOrderedArray(n,10)
	//fmt.Println(arr2)
	//sortTestHelper.TestSort("insertion sort pro",insertionSortPro,arr2)

}
