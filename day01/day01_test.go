package day01

import (
	"testing"
)

func Test_baseFuel(t *testing.T) {
	type args struct {
		mass int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Mass 12", args: args{12}, want: 2},
		{name: "Mass 14", args: args{14}, want: 2},
		{name: "Mass 1969", args: args{1969}, want: 654},
		{name: "Mass 100756", args: args{100756}, want: 33583},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := baseFuel(tt.args.mass); got != tt.want {
				t.Errorf("baseFuel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_compoundedFuel(t *testing.T) {
	type args struct {
		mass int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Mass 14", args: args{14}, want: 2},
		{name: "Mass 1969", args: args{1969}, want: 966},
		{name: "Mass 100756", args: args{100756}, want: 50346},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compoundedFuel(tt.args.mass); got != tt.want {
				t.Errorf("compoundedFuel() = %v, want %v", got, tt.want)
			}
		})
	}
}
