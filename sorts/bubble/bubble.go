package main

import (
	"fmt"
	"math/rand"
)

// BubbleSort 冒泡排序
func BubbleSort(array []int) {
	length := len(array)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			if array[j] > array[j+1] {
				temp := array[j]
				array[j] = array[j+1]
				array[j+1] = temp
			}
		}
	}
}

// MakeArray 生成数组
func MakeArray(n int, len int) []int {
	array := []int{}
	for i := 0; i < len; i++ {
		num := rand.Intn(n)
		array = append(array, num)
	}
	return array
}

func main() {
	array := []int{56, 65, 9, 16, 2, 82, 78, 31, 11, 56, -1}
	fmt.Println("before array:", array)
	BubbleSort(array)
	fmt.Println("after ", array)
}
