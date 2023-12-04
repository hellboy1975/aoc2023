package day2

import (
	"aoc2023/src/base"
	"fmt"
	"regexp"
	"strings"
)

const max_red = 12
const max_green = 13
const max_blue = 14

func GetGameId(s string) int {

	s := strings.Split(s, ":")
	return 666
}

// takes some elvish raw data, and pulls the ID out of it.
// sample: Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
func GetColourCount(s string) (blue int, red int, green int) {

	re := regexp.MustCompile("\d+:")
	r := re.MatchString()
	return 11, 12, 13
}

// Day 1, Part 1 (and now Part 2) of the AoC2023 challenge
func Part1() {
	fmt.Println("Day 2, Part 1: Cube Conundrum")

	file := base.GetDayDataFile(2, 1)

	_, err := base.ReadLines(file)
	if err != nil {
		panic(err)
	}
}

func Part2() {
	fmt.Println("Day 2, Part 2: Cube Conundrum")
}
