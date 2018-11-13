# 二分对半查找

## 介绍
二分查找法是对一组有序的数字中进行查找，传递相应的数据，进行比较查找到与原数据相同的数据，查找到了返回对应的数组下标，没有找到返回-1。


## 示意图
![binary-search](https://images2015.cnblogs.com/blog/461877/201607/461877-20160721092729169-843824718.gif)

## Code

### 非递归查找
```go
// 二分查找普通实现(非递归)  根据目标值查找出在有序数组中的位置。
// in 		有序数组
// target	查找索引的值
// index    	值的索引，不存在则返回-1
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

```

### 递归查找
```go
// 二分查找递归实现  根据目标值查找出在有序数组中的位置。
// array 	有序数组
// start	数组低下标
// end		数组高下标
// target	查找索引的值
// index    	值的索引，不存在则返回-1
func RecursiveBinarySearch(array []int,start,end int,  target int) (index int) {
	mid := (end - start) / 2 + start
	if array[mid] == target {
		return mid
	}
	if start >= end {
		return  -1
	}
	if target > array[mid] {
		return RecursiveBinarySearch(array, mid + 1, end, target)
	}
	if target < array[mid] {
		return RecursiveBinarySearch(array, start, mid - 1, target)
	}
	return -1
}
```
参考: 
* https://blog.csdn.net/lovesummerforever/article/details/78511965#commentBox
* https://github.com/liyu4/learn-algorithm-365/blob/master/search/search.go
