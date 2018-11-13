package binary_search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	type args struct {
		in     []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{1, 3, 5, 8, 10, 15}, 6}, -1},
		{"2", args{[]int{1, 3, 5, 8, 10, 15}, 1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.args.in, tt.args.target); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecursiveBinarySearch(t *testing.T) {
	type args struct {
		array  []int
		start  int
		end    int
		target int
	}
	arr1 := []int{1, 3, 5, 8, 10, 15}
	tests := []struct {
		name      string
		args      args
		wantIndex int
	}{
		// TODO: Add test cases.
		{"1", args{arr1, 0, len(arr1) - 1, 6}, -1},
		{"2", args{arr1, 0, len(arr1) - 1, 3}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIndex := RecursiveBinarySearch(tt.args.array, tt.args.start, tt.args.end, tt.args.target); gotIndex != tt.wantIndex {
				t.Errorf("RecursiveBinarySearch() = %v, want %v", gotIndex, tt.wantIndex)
			}
		})
	}
}
