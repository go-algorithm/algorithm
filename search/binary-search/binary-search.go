package binary_search

import "fmt"

// 二分查找普通实现(非递归)  根据目标值查找出在有序数组中的位置。
// in 		有序数组
// target	查找索引的值
// index    值的索引，不存在则返回-1
func BinarySearch(array []int, target int) (index int) {
	i, j := 0, len(array)-1
	for i < j {
		mid := (i + j) / 2
		fmt.Println(mid)
		if target > array[mid] {
			i = mid + 1
		} else {
			j = mid
		}
	}
	if array[i] == target {
		return i
	}
	return -1
}

// 二分查找递归实现  根据目标值查找出在有序数组中的位置。
// array 	有序数组
// start	数组低下标
// end		数组高下标
// target	查找索引的值
// index    值的索引，不存在则返回-1
func RecursiveBinarySearch(array []int, start, end int, target int) (index int) {
	mid := (end-start)/2 + start
	if array[mid] == target {
		return mid
	}
	if start >= end {
		return -1
	}
	if target > array[mid] {
		return RecursiveBinarySearch(array, mid+1, end, target)
	}
	if target < array[mid] {
		return RecursiveBinarySearch(array, start, mid-1, target)
	}
	return -1
}
