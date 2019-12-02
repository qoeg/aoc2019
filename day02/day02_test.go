package day02

import (
	"reflect"
	"testing"
)

func Test_compute(t *testing.T) {
	type args struct {
		intcode []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Test 1", args{[]int{1,0,0,0,99}}, []int{2,0,0,0,99}},
		{"Test 2", args{[]int{2,3,0,3,99}}, []int{2,3,0,6,99}},
		{"Test 3", args{[]int{2,4,4,5,99,0}}, []int{2,4,4,5,99,9801}},
		{"Test 4", args{[]int{1,1,1,4,99,5,6,0,99}}, []int{30,1,1,4,2,5,6,0,99}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compute(tt.args.intcode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("compute() = %v, want %v", got, tt.want)
			}
		})
	}
}
