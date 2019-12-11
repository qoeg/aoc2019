package day10

import (
	"math"
	"sort"
	"strconv"
	"github.com/qoeg/aoc2019/spatial"
)

type asteroid struct {
	cell spatial.Cell
	sightlines map[float64]asteroid
}

var (
	asteroids []asteroid
	best asteroid
)

func init() {
	grid := spatial.Parse(Input)
	asteroids = buildMap(grid)

	var bestCount int
	for _, a := range asteroids {
		if len(a.sightlines) > bestCount {
			best = a
			bestCount = len(a.sightlines)
		}
	}
}

func buildMap(grid [][]spatial.Cell) []asteroid {
	asteroids := []asteroid{}
	for y := 0; y < len(grid[0]); y++ {
		for x := 0; x < len(grid); x++ {
			if grid[x][y].Mark() == '#' {
				asteroids = append(asteroids, asteroid{
					cell: grid[x][y],
					sightlines: make(map[float64]asteroid),
				})
			}
		}
	}

	for i, base := range asteroids {
		for j, target := range asteroids {
			if j == i {
				continue
			}
			
			basePos := base.cell.Pos()
			targetPos := target.cell.Pos()
			sightline := math.Atan2(float64(basePos.X-targetPos.X), float64(targetPos.Y-basePos.Y))
			sightline = (sightline + math.Pi) / math.Pi * 180
			if sightline == 360 {
				sightline = 0
			}
	
			if sl, ok := base.sightlines[sightline]; ok {
				if spatial.Distance(basePos, sl.cell.Pos()) < spatial.Distance(basePos, targetPos) {
					continue
				}
			}

			base.sightlines[sightline] = target
		}
	}

	return asteroids
}

// Answer1 returns the first puzzle answer
func Answer1() string {
	return strconv.Itoa(len(best.sightlines))
}

// Answer2 returns the second puzzle answer
func Answer2() string {
	sightlines := make([]float64, len(best.sightlines))
	
	var index int
	for key := range best.sightlines {
		sightlines[index] = key
		index++
	}

	sort.Float64s(sightlines)

	lucky200 := best.sightlines[sightlines[199]].cell.Pos()

	return strconv.Itoa(lucky200.X * 100 + lucky200.Y)
}
