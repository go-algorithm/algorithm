package main

import (
	"reflect"
	"testing"
)

func Test_BubbleSort(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"t1", args{MakeArray(1000, 100)}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BubbleSort(tt.args.array)
		})
	}
}

/* 	yuyongbo-test[RAM:8G,CPU:i7-7700,OS:Win10-64bit]
Running tool: C:\Go\bin\go.exe test -timeout 30s github.com\go-algorithm\algorithm\sorts\bubble -run ^Test_BubbleSort$
ok  	github.com/go-algorithm/algorithm/sorts/bubble	0.169s
Success: Tests passed.
*/

func Benchmark_BubbleSort(b *testing.B) {
	s := MakeArray(1000, 10)
	// s := MakeArray(1000, 100)
	//s := MakeArray(1000, 1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BubbleSort(s)
	}
}

/*	yuyongbo-test[RAM:8G,CPU:i7-7700,OS:Win10-64bit]
	1000, 10	Benchmark_BubbleSort-8   	20000000	        54.0 ns/op	       0 B/op	       0 allocs/op
	1000, 100 	Benchmark_BubbleSort-8   	  200000	      5118 ns/op	       0 B/op	       0 allocs/op
	1000, 1000 	Benchmark_BubbleSort-8   	    3000	    489970 ns/op	       0 B/op	       0 allocs/op
*/

func Test_MakeArray(t *testing.T) {
	type args struct {
		n   int
		len int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "t1", args: args{100, 10}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeArray(tt.args.n, tt.args.len); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
