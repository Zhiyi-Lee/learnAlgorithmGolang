package quickSort

import (
	"learnAlgorithm/insertionSort"
	"math/rand"
	"time"
)

//对 [l,r] 进行分割
//返回 p 的位置，是的 arr[l...p-1] < arr[p] ; arr[p+1...r] > arr[p]
func __partition(arr []int, l int, r int) int {
	//v:=arr[l]
	// [l+1,j] < v   [j+1,i-1] > v
	j := l
	for i := l + 1; i <= r; i++ {
		if arr[i] < arr[l] {
			//arr[++j],arr[i]=arr[i],arr[j]  //这样不行哦 golang 不支持 ++i
			arr[j+1], arr[i] = arr[i], arr[j+1]
			j++
		}
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}

//对 [l,r] 进行分割
//返回 p 的位置，是的 arr[l...p-1] < arr[p] ; arr[p+1...r] > arr[p]
func __partitionPro(arr []int, l int, r int) int {
	//v:=arr[l]
	// [l+1,j] < v   [j+1,i-1] > v
	//随机选一个元素交换到开头 减少退化为O(n**2)的可能性
	randItem := rand.Intn(r-l+1) + l
	arr[randItem], arr[l] = arr[l], arr[randItem]
	j := l
	for i := l + 1; i <= r; i++ {
		if arr[i] < arr[l] {
			//arr[++j],arr[i]=arr[i],arr[j]  //这样不行哦 golang 不支持 ++i
			arr[j+1], arr[i] = arr[i], arr[j+1]
			j++
		}
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}

//对 [l,r] 进行快速排序
func __quickSort(arr []int, l int, r int) {
	if l >= r {
		return
	}

	p := __partition(arr, l, r)
	__quickSort(arr, l, p-1)
	__quickSort(arr, p+1, r)
}

func Sort(arr []int) {
	n := len(arr)
	__quickSort(arr, 0, n-1)
}

//对 [l,r] 进行快速排序
func __quickSortPro(arr []int, l int, r int) {
	if r-l <= 15 {
		insertionSort.SortSection(arr, l, r)
		return
	}

	p := __partitionPro(arr, l, r)
	__quickSortPro(arr, l, p-1)
	__quickSortPro(arr, p+1, r)
}

func SortPro(arr []int) {
	n := len(arr)
	rand.Seed(time.Now().UnixNano())
	__quickSortPro(arr, 0, n-1)
}

//Quick Sort 3 Ways

//处理 arr[l...r] 分为 <v ==v >v 三部分
//返回 a,b arr[a...b]==v
func __partition3Way(arr []int, l int, r int) (int, int) {
	randItem := rand.Intn(r-l+1) + l
	arr[l], arr[randItem] = arr[randItem], arr[l]

	//v := arr[l]
	//运行时 arr[l+1...lt] < v  arr[gt...r] > v  arr[lt+1...i] == v
	lt, gt := l, r+1          //保证初始下三个区间都是空的
	for i := l + 1; i < gt; { //i<=gt 如果没有大于v的元素 最后一趟会处理索引 r 的元素
		if arr[i] > arr[l] {
			arr[i], arr[gt-1] = arr[gt-1], arr[i]
			gt--
		} else if arr[i] < arr[l] {
			arr[i], arr[lt+1] = arr[lt+1], arr[i]
			lt++
			i++
		} else {
			i++
		}
	}

	//把第一个元素换到 == 的区间去
	arr[l], arr[lt] = arr[lt], arr[l]
	lt--

	return lt + 1, gt - 1
}

func __quickSort3Way(arr []int, l int, r int) {
	if r-l <= 15 {
		insertionSort.SortSection(arr, l, r)
		return
	}

	lt, gt := __partition3Way(arr, l, r) // arr[a...b] == v
	__quickSort3Way(arr, l, lt-1)
	__quickSort3Way(arr, gt+1, r)
}

func Sort3Way(arr []int) {
	n := len(arr)
	rand.Seed(time.Now().UnixNano())
	__quickSort3Way(arr, 0, n-1)
}
