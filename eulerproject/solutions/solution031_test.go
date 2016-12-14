package solutions

import "testing"
import "fmt"

//
// https://projecteuler.net/problem=31
//

func TestSolution31_01(t *testing.T) {
	count := 0
	for v200 := 0; v200 <= 200; v200 += 200 {
		for v100 := 0; v100 <= 200; v100 += 100 {
			for v050 := 0; v050 <= 200; v050 += 50 {
				for v020 := 0; v020 <= 200; v020 += 20 {
					for v010 := 0; v010 <= 200; v010 += 10 {
						for v005 := 0; v005 <= 200; v005 += 5 {
							for v002 := 0; v002 <= 200; v002 += 2 {
								for v001 := 0; v001 <= 200; v001++ {
									total := v001 + v002 + v005 + v010 + v020 + v050 + v100 + v200
									if total == 200 {
										// fmt.Println(v001, v002, v005, v010, v020, v050, v100, v200)
										count++
									}
								}
							}
						}
					}
				}
			}
		}
	}
	fmt.Println(count)
}

var coins = []int{200, 100, 50, 20, 10, 5, 2, 1}

func findposs(money int, maxcoin int) int {
	sum := 0
	if maxcoin == 7 {
		return 1
	}
	for i := maxcoin; i < 8; i++ {
		if money-coins[i] == 0 {
			sum++
		} else if money-coins[i] > 0 {
			sum += findposs(money-coins[i], i)
		}
	}
	return sum
}

func TestSolution31_02(t *testing.T) {
	fmt.Println(findposs(10, 0))

}
