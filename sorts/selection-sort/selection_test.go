package main

import (
	_ "fmt"
	"reflect"
	"testing"
)

func TestMakeArray(t *testing.T) {
	type args struct {
		scope int
		len   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeArray(tt.args.scope, tt.args.len); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectionSort(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"t1", args{MakeArray(100, 10)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SelectionSort(tt.args.array)
		})
	}
}
