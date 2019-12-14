package day12

import (
	"fmt"
	"math"
)

type moon struct {
	posx int
	posy int
	posz int
	velx int
	vely int
	velz int
}

func (m *moon) kin() int {
	return int(math.Abs(float64(m.velx)) + math.Abs(float64(m.vely)) + math.Abs(float64(m.velz)))
}

func (m *moon) pot() int {
	return int(math.Abs(float64(m.posx)) + math.Abs(float64(m.posy)) + math.Abs(float64(m.posz)))
}

func (m *moon) tot() int {
	return m.pot() * m.kin()
}

func (m *moon) grav(other *moon) {
	switch {
	case other.posx > m.posx:
		m.velx++
	case other.posx < m.posx:
		m.velx--
	}

	switch {
	case other.posy > m.posy:
		m.vely++
	case other.posy < m.posy:
		m.vely--
	}

	switch {
	case other.posz > m.posz:
		m.velz++
	case other.posz < m.posz:
		m.velz--
	}
}

func (m *moon) move() {
	m.posx += m.velx
	m.posy += m.vely
	m.posz += m.velz
}

func (m *moon) print() {
	fmt.Printf("pos=<x=%d, y=%d, z=%d>, vel=<x=%d, y=%d, z=%d>\n", m.posx, m.posy, m.posz, m.velx, m.vely, m.velz)
}

func (m *moon) sameX(other *moon) bool {
	return m.posx == other.posx && m.velx == other.velx
}

func (m *moon) sameY(other *moon) bool {
	return m.posy == other.posy && m.vely == other.vely
}

func (m *moon) sameZ(other *moon) bool {
	return m.posz == other.posz && m.velz == other.velz
}

func (m *moon) same(other *moon) bool {
	return m.sameX(other) && m.sameY(other) && m.sameZ(other)
}
