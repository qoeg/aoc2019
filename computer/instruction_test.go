package computer

import (
	"reflect"
	"testing"
)

func Test_newInstruction(t *testing.T) {
	type args struct {
		value   int
		pointer int
	}
	tests := []struct {
		name string
		args args
		want instruction
	}{
		{"Test 1", args{1002, 0}, instruction{opcode: 2, parameters: 3, parameterModes: []int{0, 1, 0}}},
		{"Test 2", args{1101, 0}, instruction{opcode: 1, parameters: 3, parameterModes: []int{1, 1, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newInstruction(tt.args.value, tt.args.pointer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newInstruction() = %v, want %v", got, tt.want)
			}
		})
	}
}
