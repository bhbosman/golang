package solutions

import (
	"bufio"
	"io"
	"os"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func readFileData(filename string) ([]string, int, error) {
	var data []string
	longestString := 0
	file, err := os.Open(filename)
	if err != nil {
		return nil, 0, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	eof := false
	for !eof {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			err = nil
			eof = true
		} else if err != nil {
			return nil, 0, err
		}
		if len(line) > 0 {
			longestString = max(longestString, len(line))
			data = append(data, string(line))
		}
	}

	return data, longestString, nil
}
