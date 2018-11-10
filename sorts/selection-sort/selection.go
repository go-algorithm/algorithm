package main

import (
	"fmt"
	"math/rand"
	"time"
)

// SelectionSort 选择排序
func SelectionSort(array []int) {
	length := len(array)
	var mIndex int
	for i := 0; i < length; i++ {
		mIndex = i
		for j := i + 1; j < length; j++ {
			if array[j] < array[mIndex] { // 寻找最小的数
				mIndex = j
			}
		}
		array[i],array[mIndex] = array[mIndex],array[i]
	}
}

// MakeArray 生成数组
func MakeArray(scope int, len int) []int {
	array := []int{}
	array = GenerateRandomNumber(0, scope, len)
	return array
}

// 生成count个[start,end)结束的不重复的随机数
func GenerateRandomNumber(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}

	//存放结果的slice
	nums := make([]int, 0, count)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn((end - start)) + start

		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			nums = append(nums, num)
		}
	}

	return nums
}

func main() {
	array := MakeArray(100, 10)
	fmt.Println("sort before array:", array)
	SelectionSort(array)
	fmt.Println("sort after array: ", array)
}
