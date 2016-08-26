package main

import (
	"fmt"

	"github.com/bhbosman/golang/eulerproject/solutions"
)

func main() {
	fmt.Print(solutions.Multiplesof3and5(1000))
	for index := 1; index <= 10; index++ {
		fmt.Println(solutions.Fibonacci(index))
	}
}
