package day1

import (
	"aoc2023/src/base"
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

const NumberMatch = `(?=(one|two|three|four|five|six|seven|eight|nine|1|2|3|4|5|6|7|8|9|0))`

// checks is the passed string is really a number.
func isNumeric(word string) (bool, int) {
	x := regexp.MustCompile(`\d`).MatchString(word)
	i := 0
	if x {
		i, _ = strconv.Atoi(word)
	}
	return x, i
}

func stringToInt(s string) (int, error) {
	switch s {
	case "one":
		return 1, nil
	case "two":
		return 2, nil
	case "three":
		return 3, nil
	case "four":
		return 4, nil
	case "five":
		return 5, nil
	case "six":
		return 6, nil
	case "seven":
		return 7, nil
	case "eight":
		return 8, nil
	case "nine":
		return 9, nil
	}

	is_num, i := isNumeric(s)
	if is_num {
		return i, nil
	}

	return -1, errors.New("no number match")
}

func findFirst(s string) int {
	// fmt.Println("findFirst", s)
	content := []byte(s)
	// this won't work, as the there are no regex libraries that support the query in NumberMatch
	re := regexp.MustCompile(NumberMatch)
	idx := re.FindAllIndex(content, -1)
	// fmt.Println(idx)
	first := idx[0]

	r := s[first[0]:first[1]]
	// fmt.Println(r)

	i, err := stringToInt(r)
	if err != nil {
		panic(err)
	}
	return i
}

func findLast(s string) int {
	fmt.Println("findLast", s)
	content := []byte(s)
	re := regexp.MustCompile(NumberMatch)
	idx := re.FindAllIndex(content, -1)
	last := idx[len(idx)-1]
	fmt.Println(idx)

	r := s[last[0]:last[1]]

	i, err := stringToInt(r)
	if err != nil {
		panic(err)
	}
	return i
}

func getNumbers(x string) string {
	reg := regexp.MustCompile("[a-zA-Z]")
	return reg.ReplaceAllString(x, "")
}

// Extracts integer based Calibration value from a raw string provided by elves
func extractCalibrationValuePart1(raw string) int {

	value := getNumbers(raw)

	result := value[0:1] + value[len(value)-1:]
	// fmt.Println(raw, result)
	i, err := strconv.Atoi(result)
	if err != nil {
		panic(err)
	}

	return i
}

// Extracts integer based Calibration value from a raw string provided by elves
func extractCalibrationValuePart2(raw string) int {
	first := findFirst(raw)
	last := findLast(raw)

	return (first * 10) + last
}

// Day 1, Part 1 (and now Part 2) of the AoC2023 challenge
func Part1() {
	fmt.Println("Day 1, Part 1: Trebuchet?!")

	file := base.GetDayDataFile(1, 1)

	lines, err := base.ReadLines(file)
	if err != nil {
		panic(err)
	}

	var sum int
	for _, s := range lines {
		x := extractCalibrationValuePart1(s)
		sum += x
	}

	fmt.Println(fmt.Sprintf(" Sum: %d", sum))

}

func Part2() {
	fmt.Println("Day 1, Part 2: Trebuchet?!")
}
