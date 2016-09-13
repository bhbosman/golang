package solutions

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
	"os"
	"strings"
	"testing"
)

//
// https://projecteuler.net/problem=18
//

// type Coordinate struct {
// 	x int
// 	y int
// }

// func (c Coordinate) sum(a, b Coordinate) Coordinate {
// 	return Coordinate{
// 		x: a.x + b.x,
// 		y: a.y + b.y}
// }

type Problem18Data01 struct {
	filename     string
	size         int
	triangleData map[Coordinate]int64
}

func (data *Problem18Data01) onBoundary(s Coordinate) bool {
	b := false
	b = b || s.x == 0
	b = b || s.y == 0
	//
	return b
}
func (data *Problem18Data01) inRange(s Coordinate) bool {
	b := true
	b = b && 0 <= s.x && s.x <= data.size-s.y
	b = b && 0 <= s.y && s.y <= data.size-s.x
	//
	return b
}

func (data *Problem18Data01) iterate(input map[Coordinate]*big.Int) map[Coordinate]*big.Int {
	result := map[Coordinate]*big.Int{}
	directions := []Coordinate{
		{
			x: 1,
			y: 0},
		{
			x: 0,
			y: 1,
		},
	}

	for c, v := range input {
		for _, d := range directions {
			s := c.sum(c, d)
			if data.inRange(s) {
				_, ok := result[s]
				if !ok {
					if data.onBoundary(s) {
						result[s] = big.NewInt(0).Add(
							big.NewInt(data.triangleData[s]),
							v)
					} else {
						result[s] = big.NewInt(0).Add(v, big.NewInt(data.triangleData[s]))
					}
				} else {
					newValue := big.NewInt(0).Add(
						big.NewInt(data.triangleData[s]),
						v)
					// fmt.Println(v, data.triangleData[s], "=", newValue, result[s], newValue.Cmp(result[s]))
					// fmt.Println(result[s])
					if newValue.Cmp(result[s]) == 1 {
						result[s] = newValue
					}
				}
			}
		}
	}
	// fmt.Println(input, result)
	return result
}

func (data *Problem18Data01) loadData() error {
	file, err := os.Open(data.filename)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	eof := false
	index := 0
	for !eof {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			err = nil
			eof = true
		} else if err != nil {
			return err
		}
		if len(line) > 0 {
			r := strings.NewReader(string(line))
			l := make([]int64, 0, 100)
			for r.Len() > 0 {
				var intValue int64
				fmt.Fscanf(r, "%d", &intValue)
				l = append(l, intValue)
			}
			y := index
			x := 0
			for _, v := range l {
				data.triangleData[Coordinate{x: x, y: y}] = v
				y--
				x++
			}
			index++
		}
	}
	return nil
}

func (data *Problem18Data01) do() *big.Int {
	list := make(map[Coordinate]*big.Int)
	list[Coordinate{x: 0, y: 0}] = big.NewInt(data.triangleData[Coordinate{x: 0, y: 0}])

	for len(list) != data.size {
		list = data.iterate(list)
		if len(list) == 0 {
			break
		}
	}
	//
	result := big.NewInt(0)
	for _, v := range list {
		if result.Cmp(v) == -1 {
			result = v
		}

	}
	return result
}

func TestSolution18_01(t *testing.T) {
	data := Problem18Data01{
		filename:     "solution018_data_01.txt",
		size:         15,
		triangleData: make(map[Coordinate]int64),
	}
	data.loadData()
	fmt.Println(data.do())
}

func TestSolution18_02(t *testing.T) {
	data := Problem18Data01{
		filename:     "solution018_data_02.txt",
		size:         4,
		triangleData: make(map[Coordinate]int64),
	}
	data.loadData()
	fmt.Println(data.do())
}
func TestSolution18_03(t *testing.T) {
	data := Problem18Data01{
		filename:     "solution067_data_01.txt",
		size:         100,
		triangleData: make(map[Coordinate]int64),
	}
	data.loadData()
	fmt.Println(data.do())
}
