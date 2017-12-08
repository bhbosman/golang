package solutions

import (
	"fmt"
	"testing"
)

//
// https://projecteuler.net/problem=14
//

type problemData struct {
	Data map[int]int
}

func (data *problemData) findCount01(v int) int {
	if v == 1 {
		return 1
	}
	_, ok := data.Data[v]
	if ok {
		return data.Data[v]
	}
	if v&1 == 0 {
		count := data.findCount01(v / 2)
		return count + 1
	}
	count := data.findCount01((v*3+1)/2) + 1
	data.Data[v] = count + 1
	return count + 1
}

func TestSolution14_01(t *testing.T) {
	data := &problemData{
		Data: make(map[int]int)}
	max := 0
	result := 0
	for i := 2; i <= 1000000; i++ {
		n := data.findCount01(i)
		if n > max {
			result = i
			max = n
		}
	}
	if result != 837799 {
		t.Fail()
	}
}

func ff(b *testing.B) {

}

func TestSolution14_02(t *testing.T) {
	fmt.Println(
		testing.Benchmark(
			func(b *testing.B) {
				for index := 1; index <= b.N; index++ {
					data := &problemData{
						Data: make(map[int]int)}
					max := 0
					result := 0
					for i := 2; i <= 1000000; i++ {
						n := data.findCount01(i)
						if n > max {
							result = i
							max = n
						}
					}
					if result != 837799 {
						t.Fail()
					}
				}
			}))
}

func TestSolution14_03(t *testing.T) {
	fmt.Println(
		testing.Benchmark(
			func(b *testing.B) {
				data := &problemData{
					Data: make(map[int]int)}
				for index := 1; index <= b.N; index++ {
					max := 0
					result := 0
					for i := 2; i <= 1000000; i++ {
						n := data.findCount01(i)
						if n > max {
							result = i
							max = n
						}
					}
					if result != 837799 {
						t.Fail()
					}
				}
			}))
}
