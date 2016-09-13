package solutions

import (
	"fmt"
	"testing"
)

//
// https://projecteuler.net/problem=15
//

type Problem15Data01 struct {
	sizeX int
	sizeY int
}

func (data *Problem15Data01) do(r int, x int, y int) int {

	if x+1 <= data.sizeX {
		r = data.do(r, x+1, y)
	}

	if y+1 <= data.sizeY {
		r = data.do(r, x, y+1)
	}

	if x == data.sizeX && y == data.sizeY {
		return r + 1
	}
	return r
}

func TestSolution15_01(t *testing.T) {

	data := Problem15Data01{
		sizeX: 3,
		sizeY: 3,
	}

	fmt.Println(data.do(0, 0, 0))
}

type Coordinate struct {
	x int
	y int
}

func (c Coordinate) sum(a, b Coordinate) Coordinate {
	return Coordinate{
		x: a.x + b.x,
		y: a.y + b.y}
}

type Problem15Data02 struct {
	sizeX int
	sizeY int
}

func (data *Problem15Data02) onBoundary(s Coordinate) bool {
	b := false
	b = b || s.x == 0
	b = b || s.y == 0
	//
	return b
}
func (data *Problem15Data02) inRange(s Coordinate) bool {
	b := true
	b = b && 0 <= s.x && s.x <= data.sizeX
	b = b && 0 <= s.y && s.y <= data.sizeY
	//
	return b
}

func (data *Problem15Data02) iterate(input map[Coordinate]int) map[Coordinate]int {
	result := map[Coordinate]int{}
	directions := []Coordinate{
		{
			x: 1,
			y: 0},
		{
			x: 0,
			y: 1}}

	for c, v := range input {
		for _, d := range directions {
			s := c.sum(c, d)

			if data.inRange(s) {
				_, ok := result[s]
				if !ok {
					if data.onBoundary(s) {
						result[s] = 1
					} else {
						result[s] = v
					}
				} else {
					result[s] = result[s] + v
				}
			}
		}
	}
	// fmt.Println(input, result)
	return result
}

func (data *Problem15Data02) do() int {
	result := 0
	list := make(map[Coordinate]int)
	list[Coordinate{x: 0, y: 0}] = 0

	for {
		list = data.iterate(list)
		if len(list) == 0 {
			break
		}
		if len(list) == 1 {
			for _, v := range list {
				result = v
			}
		}
	}

	return result
}

func TestSolution15_02(t *testing.T) {
	data := Problem15Data02{
		sizeX: 10,
		sizeY: 10,
	}
	fmt.Println(data.do())
}
