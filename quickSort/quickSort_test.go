package quickSort

import (
	"learnAlgorithm/sortTestHelper"
	"testing"
)

func TestSort(t *testing.T) {
	n := 10000000
	//n := 1000
	arr := sortTestHelper.GenerateRandomArry(n, 0, n)
	//arr := sortTestHelper.GenerateRandomArry(n, 0, 10)
	//arr := sortTestHelper.GenerateNearlyOrderedArray(n, 10)
	arr2 := make([]int, len(arr))
	arr3 := make([]int, len(arr))
	copy(arr2, arr)
	copy(arr3, arr)
	//fmt.Println("before sort:", arr2)
	//quickSort3Way(arr3)
	//fmt.Println("after sort:", arr3)
	sortTestHelper.TestSort("quick sort", Sort, arr)
	sortTestHelper.TestSort("quick sort pro", SortPro, arr2)
	sortTestHelper.TestSort("quick sort 3 way", Sort3Way, arr3)
}
