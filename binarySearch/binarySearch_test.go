package binarySearch

import (
	"fmt"
	"learnAlgorithm/quickSort"
	"learnAlgorithm/sortTestHelper"
	"testing"
)

func TestSearch(t *testing.T) {
	n := 10000000
	arr := sortTestHelper.GenerateRandomArry(n, 0, n)
	quickSort.SortPro(arr)
	//fmt.Println(arr)
	targetIndex := Search(arr, 10000)
	if targetIndex > -1 {
		fmt.Println(targetIndex, " ", arr[targetIndex])
	} else {
		fmt.Println("no find")
	}
}
