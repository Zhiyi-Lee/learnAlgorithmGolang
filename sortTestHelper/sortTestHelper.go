package sortTestHelper

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

//生成一个长度为 n 范围 [rangeL,ranger] 的切片
func GenerateRandomArry(n int, rangeL int, rangeR int) []int {
	var arr []int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Intn(rangeR-rangeL+1)+rangeL)
		//arr = append(arr, rand.Int() % (rangeR - rangeL +1) +rangeL)
	}
	return arr
}

//生成一个近乎有序的数组切片
func GenerateNearlyOrderedArray(n int, swapTimes int) []int {
	var arr []int
	for i := 0;i<n;i++ {
		arr = append(arr,i)
	}

	//rand swap
	for i := 0; i < swapTimes; i++ {
		posx:=rand.Int()%n
		posy:=rand.Int()%n
		arr[posx],arr[posy]=arr[posy],arr[posx]
	}

	return arr
}

func TestSort(sortName string, sort func(arr []int), arr []int) {
	startTime := time.Now()
	sort(arr)
	spendTime := time.Since(startTime)

	if !isSorted(arr) {
		log.Fatal("sort fail.")
	}

	fmt.Println(sortName, ":","n=",len(arr)," spend time:", spendTime)
}

func isSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}

	return true
}

//获取两个int的最小值
func Min(a int, b int) int {
	if a<b {
		return a
	}
	return b
}