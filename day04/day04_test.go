package day04

import (
	"reflect"
	"testing"
)

func Test_digits6(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Test 1", args{123456}, []int{1, 2, 3, 4, 5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := digits6(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("digits6() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pairsOnly(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Test 1", args{[]int{5,7,8,8,8,9}}, false},
		{"Test 2", args{[]int{5,7,8,9,9,9}}, false},
		{"Test 3", args{[]int{5,7,8,8,9,9}}, true},
		{"Test 4", args{[]int{5,8,9,9,9,9}}, false},
		{"Test 5", args{[]int{5,8,8,8,8,9}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pairsOnly(tt.args.digits); got != tt.want {
				t.Errorf("pairsOnly() = %v, want %v", got, tt.want)
			}
		})
	}
}
