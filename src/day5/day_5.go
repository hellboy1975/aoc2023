package day5

import (
	"aoc2023/src/base"
	"fmt"
	"strings"
)

var file string

type mapRange struct {
	low  int
	high int
}

func init() {
	file = base.GetDayDataFile("5", "1")
}

// seeds are in the first line
func ParseSeeds(line string) []int {

	return base.StringToIntArray(strings.Split(line, ":")[1])
}

func Part1() {
	fmt.Println("Day 5, Part 1: If You Give A Seed A Fertilizer")

	lines, err := base.ReadLines(file)
	if err != nil {
		panic(err)
	}

	seeds := ParseSeeds(lines[0])
	fmt.Println("seeds:", seeds)
}

func Part2() {
	fmt.Println("Day 5, Part 2: If You Give A Seed A Fertilizer")
}
