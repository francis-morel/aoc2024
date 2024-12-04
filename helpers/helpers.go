package helpers

import (
	"bufio"
	"os"
)

func ReadFile(file string, callback func(string)) {
	readFile, err := os.Open(file)
	if err != nil {
		panic("Failed to read the file")
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		callback(fileScanner.Text())
	}
}

func Abs(value int) int {
	if value < 0 {
		return -value
	} else {
		return value
	}
}

func ReadWholeFile(file string) []byte {
	data, err := os.ReadFile(file)

	if err != nil {
		panic(err)
	}

	return data
}
