package solutions

import (
	"fmt"
	"sort"
	"strconv"
	"testing"
)

func TestSolution24_01(t *testing.T) {
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	extractKeys := func(data map[int]int) []int {
		result := make([]int, 0, len(data))
		for key := range data {
			result = append(result, key)
		}
		return result
	}
	concatarrays := func(a []int, b []int) []int {
		result := make([]int, 0, len(a)+len(b))
		for _, value := range a {
			result = append(result, value)
		}
		for _, value := range b {
			result = append(result, value)
		}
		return result
	}
	print := func(data []int) string {
		result := ""
		for _, value := range data {
			result = result + strconv.Itoa(value)
		}
		return result

	}
	counter := 0
	for {
		counter++

		if counter > 1000010 {
			break
		}

		if counter == 1000000 {
			fmt.Println(counter, data, print(data))
		}

		beginIndex := len(data) - 1

		found := false
		for i := len(data) - 1; i >= 1; i-- {
			if data[i] > data[i-1] {
				beginIndex = i
				found = true
				break
			}
		}
		if found {
			dealWithBlock := func(data []int) ([]int, bool) {
				findNextNumberInBlock := func(data []int) (int, bool) {
					tree := make(map[int]int)
					nextValue := data[0] + 1
					nextInTreeFound := false
					var nextInTree int
					for _, value := range data {
						if value >= nextValue {
							tree[value] = 0
							if !nextInTreeFound {
								nextInTreeFound = true
								nextInTree = value
							} else if value < nextInTree {
								nextInTree = value
							}
						}
					}
					return nextInTree, nextInTreeFound
				}

				if nextInTree, ok := findNextNumberInBlock(data); ok {
					tree := make(map[int]int)
					for _, value := range data {
						tree[value] = 0
					}

					result := make([]int, 1, 1)
					result[0] = nextInTree

					delete(tree, nextInTree)
					a := extractKeys(tree)
					sort.Ints(a)
					result = concatarrays(result, a)
					return result, true
				}
				return nil, false
			}
			for {
				block := data[beginIndex-1:]
				block, exit := dealWithBlock(block)
				if exit {
					data = concatarrays(data[:beginIndex-1], block)
					break
				}
				beginIndex--
			}

		}
	}
}
func fact(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}
func TestSolution24_02(t *testing.T) {

	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	N := len(data)
	M := 1000000
	sol := make([]int, 0, len(data))
	fmt.Println(data, N, M, sol)
	for k := 1; k <= N; k++ {
		idx := (M - 1) / fact(N-k)
		sol = append(sol, data[idx])
		data = RemoveIndex(data, idx)
		M = M - idx*fact(N-k)
	}
	print := func(data []int) string {
		result := ""
		for _, value := range data {
			result = result + strconv.Itoa(value)
		}
		return result

	}
	if "2783915460" != print(sol) {
		t.Fatal("ddddd")
	}
	fmt.Println(print(sol))
}
