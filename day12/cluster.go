package day12

type cluster []*moon

func (c cluster) clone() cluster {
	result := make([]*moon, len(c))
	for i := range c {
		m := &moon{}
		*m = *c[i]
		result[i] = m
	}
	return result
}

func (c cluster) energy() int {
	sum := 0
	for _, m := range []*moon(c) {
		sum += m.tot()
	}
	return sum
}

func (c cluster) step() {
	for _, m := range []*moon(c) {
		for _, n := range []*moon(c) {
			if m == n {
				continue
			}

			m.grav(n)
		}
	}

	for _, m := range []*moon(c) {
		m.move()
	}
}

func (c cluster) print() {
	for _, m := range []*moon(c) {
		m.print()
	}
}

func sameX(c, other cluster) bool {
	result := true
	for i := 0; i < len(c); i++ {
		result = result && []*moon(c)[i].sameX([]*moon(other)[i])
	}
	return result
}

func sameY(c, other cluster) bool {
	result := true
	for i := 0; i < len(c); i++ {
		result = result && []*moon(c)[i].sameY([]*moon(other)[i])
	}
	return result
}

func sameZ(c, other cluster) bool {
	result := true
	for i := 0; i < len(c); i++ {
		result = result && []*moon(c)[i].sameZ([]*moon(other)[i])
	}
	return result
}

func same(c, other cluster) bool {
	result := true
	for i := 0; i < len(c); i++ {
		result = result && []*moon(c)[i].same([]*moon(other)[i])
	}
	return result
}
