package main

import(
	"fmt"
	"testing"
	"math/rand"
)
func TestSort(t *testing.T){
	fmt.Println("单元测试....start!!!")
	array:=[]int{}
	for i:=0;i<10;i++{
		num:=rand.Intn(500)
		array=append(array,num)
	}
	fmt.Println("sort array before is ",array)
	sort(array)
	fmt.Println("sort array after is ",array)
}