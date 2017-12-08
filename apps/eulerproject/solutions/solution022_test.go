package solutions

import (
	"encoding/csv"
	"os"
	"sort"
	"testing"
)

//
// https://projecteuler.net/problem=22
//

func TestSolution22_01(t *testing.T) {
	file, err := os.Open("solution022_names.txt")
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	data, err := reader.Read()
	if err != nil {
		t.Error(err)
	}
	sort.Strings(data)

	NameScore := func(name string) int {
		sum := 0
		for _, v := range name {
			sum += int(1 + v - 'A')
		}
		return sum
	}
	sum := 0
	for i, v := range data {
		sum += ((i + 1) * NameScore(v))
	}
	if 871198282 != sum {
		t.Fail()
	}
}
