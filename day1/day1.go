package day1

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Run() {
	RunPart1()
	RunPart2()
}

func RunPart1() {
	left, right := getValues()

	slices.Sort(left[:])
	slices.Sort(right[:])

	result := 0

	for i := range 1000 {
		result += abs(left[i] - right[i])
	}

	fmt.Print("Part 1: ")
	fmt.Println(result)
}

func RunPart2() {
	left, right := getValues()

	mapped := make(map[int]int)

	for _, val := range left {
		mapped[val] = 0
	}

	for _, val := range right {
		_, ok := mapped[val]
		if !ok {
			continue
		}

		mapped[val]++
	}

	result := 0

	for key, val := range mapped {
		result += key * val
	}

	fmt.Println("Part 2:", result)
}

func abs(value int) int {
	if value < 0 {
		return -value
	} else {
		return value
	}
}

func getValues() (*[1000]int, *[1000]int) {
	readFile, err := os.Open("day1/input.txt")
	if err != nil {
		panic("Failed to read the file")
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var left [1000]int
	var right [1000]int

	currentIndex := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		split := strings.Split(line, "   ")
		left[currentIndex], _ = strconv.Atoi(split[0])
		right[currentIndex], _ = strconv.Atoi(split[1])
		currentIndex++
	}

	return &left, &right
}
