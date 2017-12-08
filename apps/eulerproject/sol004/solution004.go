package sol004

import (
	"fmt"
	"sync"
)

func FindPalindrome() int {
	inputProductStream := &InputSumStream{
		wg:    sync.WaitGroup{},
		input: make(chan ProductInput, 1024)}

	outputProductStream := &OutputSumStream{
		wg:    sync.WaitGroup{},
		input: make(chan ProductOutput, 1024)}

	palindromeStream := &PalindromeStream{
		wg:    sync.WaitGroup{},
		input: make(chan int, 1024)}

	aa := make(chan int)

	go processProduct(inputProductStream, outputProductStream, 1)
	go processProduct(inputProductStream, outputProductStream, 2)
	go processProduct(inputProductStream, outputProductStream, 3)
	go processProduct(inputProductStream, outputProductStream, 4)

	go processFindPalindromes(outputProductStream, palindromeStream)
	go processFindPalindromes(outputProductStream, palindromeStream)
	go processFindPalindromes(outputProductStream, palindromeStream)
	go processFindPalindromes(outputProductStream, palindromeStream)

	go processFindHighestPalindromes(palindromeStream, aa)

	beginNumber := 100
	endNumber := 999

	for x := beginNumber; x <= endNumber; x++ {
		for y := x; y <= endNumber; y++ {
			input := ProductInput{x: x, y: y}
			inputProductStream.add(input)
		}
	}
	inputProductStream.wg.Wait()
	inputProductStream.close()

	outputProductStream.wg.Wait()
	outputProductStream.close()

	palindromeStream.wg.Wait()
	palindromeStream.close()

	return <-aa
}

func processProduct(input *InputSumStream, output *OutputSumStream, number int) {
	for v := range input.input {
		func() {
			defer input.wg.Done()
			//
			ans := ProductOutput{
				product: v.x * v.y,
				count:   number}
			output.add(ans)
		}()
	}
}

func processFindPalindromes(input *OutputSumStream, output *PalindromeStream) {
	for v := range input.input {
		func() {
			defer input.wg.Done()

			if isPalindrome(fmt.Sprintf("%d", v.product)) {
				output.add(v.product)
			}
		}()
	}
}

func isPalindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}

func processFindHighestPalindromes(input *PalindromeStream, output chan<- int) {
	result := 0
	for v := range input.input {
		func() {
			defer input.wg.Done()
			if v > result {
				result = v
			}
		}()
	}
	output <- result
}
