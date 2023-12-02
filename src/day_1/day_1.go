package day_1

import (
	"aoc2023/src/base"
	"fmt"
	"regexp"
	"strconv"
)

func getNumbers(x string) string {
	reg := regexp.MustCompile("[a-zA-Z]")
	return reg.ReplaceAllString(x, "")
}

// Extracts the Calibration value from a raw string provided by elves
func extractCalibrationValue(raw string) int {
	value := getNumbers(raw)
	result := value[0:1] + value[len(value)-1:]
	fmt.Println(value[0:1], value[len(value)-1:], result)

	i, err := strconv.Atoi(result)
	if err != nil {
		panic(err)
	}
	return i
}

// Day 1, Part 1 of the AoC2023 challenge
func Part1() {
	fmt.Println("Day 1, Part 1: Trebuchet?!")

	file := base.GetDayDataFile(1, 1)

	lines, err := base.ReadLines(file)
	if err != nil {
		panic(err)
	}

	var sum int
	for i, s := range lines {

		x := extractCalibrationValue(s)
		fmt.Println(i, s, x)
		sum += x
	}

	fmt.Println(fmt.Sprintf(" Sum: %d", sum))

}

// Day 1, Part 2 of the AoC2023 challenge
func Part2() {
	fmt.Println("Day 1, Part 2: Trebuchet?!")

}
