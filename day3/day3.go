package day3

import (
	"fmt"
	"strconv"

	"github.com/francis-morel/aoc2024/helpers"
)

func Run() {
	part1()
	part2()
}

type Numbers struct {
	first   []byte
	second  []byte
	current int
}

func (n *Numbers) append(c byte) {
	if n.current == 0 {
		n.first = append(n.first, c)
	} else {
		n.second = append(n.second, c)
	}
}

func (n *Numbers) sep() error {
	n.current++

	if n.current > 1 {
		return fmt.Errorf("Problem while parsing, too many numbers found")
	}

	return nil
}

func (n *Numbers) reset() {
	n.first = make([]byte, 0)
	n.second = make([]byte, 0)
	n.current = 0
}

func (n *Numbers) mul() (int, error) {
	if n.current != 1 {
		return 0, fmt.Errorf("Multiplying without reaching the second number")
	} else if len(n.first) < 1 {
		return 0, fmt.Errorf("Multiplying with the first number empty")
	} else if len(n.second) < 1 {
		return 0, fmt.Errorf("Multiplying with the second number empty")
	}

	result := int(parseNumber(n.first)) * int(parseNumber(n.second))

	return result, nil
}

func part1() {
	content := helpers.ReadWholeFile("day3/input.txt")

	n := new(Numbers)
	isBuildingNumbers := false
	total := 0

	for i, c := range content {
		if i < 4 {
			continue
		}

		if isBuildingNumbers {
			if isNumber(c) {
				n.append(c)
				continue
			} else if isSeparator(c) {
				if n.sep() != nil {
					n.reset()
				}
				continue
			} else if isCloseParen(c) {
				res, err := n.mul()

				if err == nil {
					total += res
				}

				n.reset()
				isBuildingNumbers = false
				continue
			} else {
				n.reset()
				isBuildingNumbers = false
				continue
			}
		}

		if isOpenParen(c) && isMulWord(content[i-3:i+1]) {
			isBuildingNumbers = true
		}
	}

	fmt.Println("Completed part 1. The result is", total)
}

func part2() {
	content := helpers.ReadWholeFile("day3/input.txt")

	n := new(Numbers)
	isBuildingNumbers := false
	total := 0
	canDo := true

	for i, c := range content {
		if i < 4 {
			continue
		}

		if isBuildingNumbers {
			if isNumber(c) {
				n.append(c)
				continue
			} else if isSeparator(c) {
				if n.sep() != nil {
					n.reset()
				}
				continue
			} else if isCloseParen(c) {
				res, err := n.mul()

				if err == nil {
					total += res
				}

				n.reset()
				isBuildingNumbers = false
				continue
			} else {
				n.reset()
				isBuildingNumbers = false
				continue
			}
		}

		if isOpenParen(c) {
			if len(content) > i+2 && isDoWord(content[i-2:i+2]) {
				canDo = true
				continue
			}

			if i > 5 && len(content) > i+2 && isDontWord(content[i-5:i+2]) {
				canDo = false
				continue
			}

			if canDo && isMulWord(content[i-3:i+1]) {
				isBuildingNumbers = true
			}
		}
	}

	fmt.Println("Completed part 2. The result is", total)
}

func isNumber(c byte) bool {
	return c >= '0' && c <= '9'
}

func isMulWord(data []byte) bool {
	return string(data) == "mul("
}

func isDoWord(data []byte) bool {
	return string(data) == "do()"
}

func isDontWord(data []byte) bool {
	return string(data) == "don't()"
}

func isSeparator(c byte) bool {
	return c == ','
}

func isValidCharacter(c byte) bool {
	return c == 'm' || c == 'a' || c == 'x'
}

func isOpenParen(c byte) bool {
	return c == '('
}

func isCloseParen(c byte) bool {
	return c == ')'
}

func parseNumber(data []byte) int {
	res, err := strconv.Atoi(string(data))

	if err != nil {
		panic(err)
	}

	return res
}
