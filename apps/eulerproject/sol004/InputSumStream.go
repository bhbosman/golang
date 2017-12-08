package sol004

import "sync"

type InputSumStream struct {
	wg    sync.WaitGroup
	input chan ProductInput
}

func (inputStream *InputSumStream) add(input ProductInput) {
	inputStream.wg.Add(1)
	inputStream.input <- input
}

func (inputStream *InputSumStream) close() {
	close(inputStream.input)
}

type PalindromeStream struct {
	wg    sync.WaitGroup
	input chan int
}

func (inputStream *PalindromeStream) add(input int) {
	inputStream.wg.Add(1)
	inputStream.input <- input
}

func (inputStream *PalindromeStream) close() {
	close(inputStream.input)
}
