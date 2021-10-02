package binarySearch

// 二分查找法在 arr 中搜索 target
// 返回元素的索引，没找到则返回 -1
//O(logN)
func Search(arr []int, target int) int {
	l := 0
	r := len(arr) - 1
	//在 [l,r] 中搜素 targer
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			//下一轮在 [l,mid-1] 中搜索target
			r = mid - 1
		} else {
			//下一轮在 [mid+1,r] 中搜索target
			l = mid + 1
		}
	}

	return -1
}
