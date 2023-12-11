package day3

import (
	"aoc2023/src/base"
	"fmt"
	"regexp"
	"unicode"
)

const CharNum = 1
const CharSymbol = 2
const CharBlank = 3

const NumRegex = "([0-9])\\w+"

// returns 1 for number, 2 for symbol and 3 for .
func CharType(c rune) int {
	if c == '.' {
		return CharNum
	} else if unicode.IsNumber(c) {
		return CharSymbol
	}
	return CharBlank
}

func LineNumberChunks(s string) [][]int {
	line := []byte(s)
	// this will return a bunch of slices which will correspond to the numbers
	x := regexp.MustCompile(NumRegex).FindAllIndex(line, -1)

	return x
}

// Creates a 2D array of boolean values indicating if a x,y coordinate is a symbol
func Symbols(lines []string) [][]bool {

	var symbols [][]bool
	for y, line := range lines {
		for x, cell := range line {
			if CharType(rune(cell)) == CharSymbol {
				symbols[x][y] = true
			} else {
				symbols[x][y] = false
			}
		}
	}
	fmt.Println()
	return symbols
}

// Takes a chunk (two x coords) and a line (the y coord)
/*

..........
...345....
..........
*/
func CalcRange(chunk []int, line int) (xcoord, ycoord []int) {

	// if the line is the first one (0) the should us that instead of -1
	if line == 0 {
		xcoord = append(xcoord, 0)
	} else {
		xcoord = append(xcoord, line-1)
	}

	// if the first array item in the chunk is 0 then the xcoord should also be 0
	if chunk[0] == 0 {
		xcoord = append(xcoord, 0)
	} else {
		xcoord = append(xcoord, chunk[0]-1)
	}

	return xcoord, ycoord
}

// Day 3, Part 1 of the AoC2023 challenge
func Part1() {
	fmt.Println("Day 3, Part 1: Gear Ratios")

	file := base.GetDayDataFile(3, 1)

	lines, err := base.ReadLines(file)
	if err != nil {
		panic(err)
	}

	// symbols := Symbols(lines)

	for x, line := range lines {
		chunks := LineNumberChunks(line)

		for c, chunk := range chunks {
			fmt.Println(x, c, chunk)
		}
	}
}

// Day 3, Part 2 of the AoC2023 challenge
func Part2() {
	fmt.Println("Day 3, Part 2: Gear Ratios")
}
