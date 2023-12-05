package day2

import (
	"aoc2023/src/base"
	"fmt"
	"strconv"
	"strings"
)

const max_red = 12
const max_green = 13
const max_blue = 14

func GetGameId(s string) int {

	id := strings.Split(s, ":")[0]
	fmt.Println(id)

	id = strings.Split(id, " ")[1]
	fmt.Println(id)

	intId, _ := strconv.Atoi(id)
	return intId
}

// takes some elvish raw data, and pulls the ID out of it.
// sample: Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
func GetColourCount(s string) (blue int, red int, green int) {

	colours := strings.Split(s, ":")[1]
	rounds := strings.Split(colours, ";")

	for _, r := range rounds {
		sets := strings.Split(r, ",")
		for i, s := range sets {
			s = strings.TrimSpace(s)
			fmt.Println("set:", s)
			colour := strings.Split(s, " ")[1]
			count, _ := strconv.Atoi(strings.Split(s, " ")[0])

			fmt.Println(i, "colour", colour)
			fmt.Println(i, "count", count)

			if colour == "blue" {
				blue += count
			} else if colour == "red" {
				red += count
			} else if colour == "green" {
				green += count
			}
		}
	}

	return blue, red, green
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
