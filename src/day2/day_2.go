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

	id = strings.Split(id, " ")[1]

	intId, _ := strconv.Atoi(id)
	return intId
}

func IsGamePossible(s string) bool {
	colours := strings.Split(s, ":")[1]
	rounds := strings.Split(colours, ";")

	for _, r := range rounds {
		sets := strings.Split(r, ",")
		for _, s := range sets {
			s = strings.TrimSpace(s)
			colour := strings.Split(s, " ")[1]
			count, _ := strconv.Atoi(strings.Split(s, " ")[0])

			if colour == "blue" {
				if count > max_blue {
					return false
				}
			} else if colour == "red" {
				if count > max_red {
					return false
				}
			} else if colour == "green" {
				if count > max_green {
					return false
				}
			}
		}
	}
	return true

}

// Works out the highest value of each colour in a given round
func GetColourMax(s string) (blue int, red int, green int) {

	colours := strings.Split(s, ":")[1]
	rounds := strings.Split(colours, ";")

	for _, r := range rounds {
		sets := strings.Split(r, ",")
		for _, s := range sets {
			s = strings.TrimSpace(s)
			colour := strings.Split(s, " ")[1]
			count, _ := strconv.Atoi(strings.Split(s, " ")[0])

			if colour == "blue" && count > blue {
				blue = count
			} else if colour == "red" && count > red {
				red = count
			} else if colour == "green" && count > green {
				green = count
			}
		}
	}

	return blue, red, green
}

// Day 2, Part 1 of the AoC2023 challenge
func Part1() {
	fmt.Println("Day 2, Part 1: Cube Conundrum")

	file := base.GetDayDataFile("2", "1")

	lines, err := base.ReadLines(file)
	if err != nil {
		panic(err)
	}

	var sum int
	for _, s := range lines {
		id := GetGameId(s)

		if IsGamePossible(s) {
			sum += id
		}
	}

	fmt.Println("Sum of possible games:", sum)
}

func Part2() {
	fmt.Println("Day 2, Part 2: Cube Conundrum")

	file := base.GetDayDataFile("2", "1")

	lines, err := base.ReadLines(file)
	if err != nil {
		panic(err)
	}

	var sum int
	for _, s := range lines {

		blue, red, green := GetColourMax(s)

		// will one of these numbers ever be zero?  Didn't turn out to be a problem...
		sum += blue * red * green
	}

	fmt.Println("Sum of possible games:", sum)
}
