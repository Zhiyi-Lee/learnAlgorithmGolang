package selectionSort

import (
	"learnAlgorithm/sortTestHelper"
	"testing"
)

func TestSort(t *testing.T) {
	arr := sortTestHelper.GenerateRandomArry(10000, 0, 10000)
	//fmt.Println("before sort:")
	//fmt.Println(arr)
	//selectionSort(arr)
	//fmt.Println("after sort:")
	sortTestHelper.TestSort("selection sort n=10000", Sort, arr)

	arr2 := sortTestHelper.GenerateRandomArry(100000, 0, 100000)
	sortTestHelper.TestSort("selcetion sort n=100000", Sort, arr2)
	//fmt.Println(arr)
}
