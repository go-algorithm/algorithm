package main
 
import (
    "fmt"
)
func sort(array []int){
	
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

func main() {
    array := []int{56,65,9,16,2,82,78,31,11,56,2}
    fmt.Println("before ",array)
	sort(array)
    fmt.Println("after ",array)
}