# 冒泡排序（Bubble Sort）

冒泡排序是一种简单的排序算法。它重复地走访过要排序的数列，一次比较两个元素，如果它们的顺序错误就把它们交换过来。走访数列的工作是重复地进行直到没有再需要交换，也就是说该数列已经排序完成。这个算法的名字由来是因为越小的元素会经由交换慢慢“浮”到数列的顶端。 

## 算法描述

1.比较相邻的元素。如果第一个比第二个大，就交换它们两个；

2.对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对，这样在最后的元素应该会是最大的数；

3.针对所有的元素重复以上的步骤，除了最后一个；

重复步骤1~3，直到排序完成。

## 动图演示

![Bubble Sort](https://images2017.cnblogs.com/blog/849589/201710/849589-20171015223238449-2146169197.gif)

## 代码示例

> bubble.go

```go
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

```

> bubble_test.go

```go
func Test_sort(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name:"array1",args: struct{ array []int }{array: []int{56,65,9,16,2,82,78,31,11,56,2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort(tt.args.array)
		})
	}
}
```


