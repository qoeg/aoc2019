package day11

type direction int

const (
	up direction = iota
	right
	down
	left
)

func (d direction) turn(change int) direction {
	if change == 0 {
		if d--; d < up {
			d = left
		}
	} else {
		if d++; d > left {
			d = up
		}
	}
	return d
}