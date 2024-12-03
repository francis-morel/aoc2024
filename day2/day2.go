package day2

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/francis-morel/aoc2024/helpers"
)

func Run() {
	part1()
	part2()
}

func part1() {
	valid := 0

	helpers.ReadFile("day2/input.txt", func(s string) {
		report := toIntArray(s)

		if validateReport(report) {
			valid++
		}
	})

	fmt.Println("Day 2 part 1:", valid)
}

func part2() {
	valid := 0

	helpers.ReadFile("day2/input.txt", func(s string) {
		report := toIntArray(s)

		if validateReport(report) || tryDampenReport(report) {
			valid++
		}
	})

	fmt.Println("Day 2 part 2:", valid)
}

func tryDampenReport(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		cp := make([]int, len(arr))
		copy(cp, arr)

		s := slices.Delete(cp, i, i+1)
		if validateReport(s) {
			return true
		}
	}

	return false
}

func toIntArray(s string) []int {
	arr := make([]int, 0)

	for _, s := range strings.Split(s, " ") {
		c, err := strconv.Atoi(s)
		if err != nil {
			panic("Failed to parse a number")
		}

		arr = append(arr, c)
	}

	return arr
}

func validateReport(arr []int) bool {
	validators := []func([]int) bool{
		ruleSameDirection,
		ruleGradually,
	}

	for _, v := range validators {
		if !v(arr) {
			return false
		}
	}

	return true
}

func ruleSameDirection(arr []int) bool {
	var direction bool

	for i := 0; i < len(arr)-1; i++ {
		curr := getDirection(arr[i], arr[i+1])
		if i != 0 && direction != curr {
			return false
		}

		direction = curr
	}

	return true
}

func getDirection(a int, b int) bool {
	// 0 is decreasing
	// 1 is increasing

	return a < b
}

func ruleGradually(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		gap := helpers.Abs(arr[i] - arr[i+1])
		if gap == 0 || gap > 3 {
			return false
		}
	}

	return true
}
