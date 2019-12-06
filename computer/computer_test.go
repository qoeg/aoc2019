package computer

import (
	"reflect"
	"testing"
)

func TestRun(t *testing.T) {
	type args struct {
		intcode []int
		input   int
		verbose bool
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 []int
	}{
		{"Test 1", args{[]int{1002,4,3,4,33}, 0, false}, 0, []int{1002,4,3,4,99}},
		{"Test 2", args{[]int{1101,100,-1,4,0}, 0, false}, 0, []int{1101,100,-1,4,99}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Run(tt.args.intcode, tt.args.input, tt.args.verbose)
			if got != tt.want {
				t.Errorf("Run() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Run() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
