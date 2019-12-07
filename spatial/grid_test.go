package spatial

import (
	"reflect"
	"testing"
)

var input = 
`#######
#E..G.#
#...#.#
#.G.#G#
#######`

func TestGrid_Neighbors(t *testing.T) {
	grid := Parse(input)

	type args struct {
		c       Coordinate
		include []rune
	}
	tests := []struct {
		name string
		g    Grid
		args args
		want []Cell
	}{
		{
			name: "Corner",
			g: grid,
			args: args{Coordinate{X:0, Y:0}, []rune{'.'}},
			want: []Cell{},
		},
		{
			name: "Partial",
			g: grid,
			args: args{Coordinate{X:4, Y:1}, []rune{'.'}},
			want: []Cell{
				NewCell('.', Coordinate{X:3, Y:1}),
				NewCell('.', Coordinate{X:5, Y:1}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.Neighbors(tt.args.c, tt.args.include...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Grid.Neighbors() = %v, want %v", got, tt.want)
			}
		})
	}
}
