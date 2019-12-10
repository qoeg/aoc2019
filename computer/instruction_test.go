package computer

import (
	"reflect"
	"testing"
)

func Test_newInstruction(t *testing.T) {
	p := New(0, Config{})
	
	type args struct {
		program *Program
		value   int64
	}
	tests := []struct {
		name string
		args args
		want instruction
	}{
		{"Test 1", args{p, 1002}, instruction{opcode: 2, parameters: 3, parameterModes: []int{0, 1, 0}}},
		{"Test 2", args{p, 1101}, instruction{opcode: 1, parameters: 3, parameterModes: []int{1, 1, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newInstruction(tt.args.program, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newInstruction() = %v, want %v", got, tt.want)
			}
		})
	}
}
