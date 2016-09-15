package solutions

import (
	"fmt"
	"testing"
)

func TestSolution19_01(t *testing.T) {
	leapYear := func(year int) bool {
		if !(year%100 == 0 && year%400 == 0) {
			return false
		}
		if year%4 == 0 {
			return true
		}
		return false
	}
	daysInYearMonth := func(year int, month int) int {
		NormalYearDays := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
		LeapYearDays := [12]int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
		if leapYear(year) {
			return LeapYearDays[month]
		}
		return NormalYearDays[month]
	}
	days := 0
	count := 0
	for year := 1901; year <= 2000; year++ {
		for month := 0; month < 12; month++ {
			count += daysInYearMonth(year, month)
			if count%7 == 6 {
				days++
			}
		}
	}
	fmt.Println(count, days)
}
