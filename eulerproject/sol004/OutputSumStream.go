package sol004

import "sync"

type OutputSumStream struct {
	wg    sync.WaitGroup
	input chan ProductOutput
}

func (inputStream *OutputSumStream) add(input ProductOutput) {
	inputStream.wg.Add(1)
	inputStream.input <- input
}

func (inputStream *OutputSumStream) close() {
	close(inputStream.input)
}
