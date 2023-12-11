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

const DataExtent = 139 // 140 columns/rows in the data provided

const NumRegex = "([0-9])\\w+"

type coord struct {
	x int
	y int
}

// returns 1 for number, 2 for symbol and 3 for .
func CharType(c rune) int {
	if c == '.' {
		return CharBlank
	} else if unicode.IsNumber(c) {
		return CharNum
	}
	return CharSymbol
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
	// symbols := make([][]bool, 139)

	fmt.Println("Lines:", len(lines))
	for y, line := range lines {
		fmt.Println("Cells:", len(line))
		fmt.Println(line)
		sline := make([]bool, 140)
		for x, cell := range line {
			// fmt.Println(x, y, cell)
			if CharType(rune(cell)) == CharSymbol {
				// fmt.Println("true")
				sline[x] = true
			} else {
				// fmt.Println("false")
				sline[x] = false
			}
		}
		fmt.Println("sline", sline)
		symbols[y] = sline
	}
	fmt.Println()
	return symbols
}

func Ycoord(chunk []int, line int) (ycoord coord) {
	// ycoord will contain the bottom-right corner of the range "box"
	if chunk[1] >= DataExtent {
		ycoord.x = DataExtent
	} else {
		ycoord.x = chunk[1] + 1
	}

	// if we're on the final line then we won't add one
	if line >= DataExtent {
		ycoord.y = DataExtent
	} else {
		ycoord.y = line + 1
	}

	return ycoord
}

func Xcoord(chunk []int, line int) (xcoord coord) {
	// xcoord will contin the top-left corner of the range "box"
	// if the line is the first one (1) the should us that instead of -1
	if line == 0 {
		xcoord.y = 0
	} else {
		xcoord.y = line - 1
	}

	// if the first array item in the chunk is 0 then the xcoord should also be 0
	if chunk[0] == 0 {
		xcoord.x = 0
	} else {
		xcoord.x = chunk[0] - 1
	}

	return xcoord
}

// Takes a chunk (two x coords) and a line (the y coord)
/*

..........
...345....
..........
*/
func CalcRange(chunk []int, line int) (xcoord, ycoord coord) {

	return Xcoord(chunk, line), Ycoord(chunk, line)
}

// Day 3, Part 1 of the AoC2023 challenge
func Part1() {
	fmt.Println("Day 3, Part 1: Gear Ratios")

	file := base.GetDayDataFile(3, 1)

	lines, err := base.ReadLines(file)
	if err != nil {
		panic(err)
	}

	symbols := Symbols(lines)
	fmt.Println(symbols)

	for x, line := range lines {
		chunks := LineNumberChunks(line)

		for c, chunk := range chunks {
			fmt.Println(x, c, chunk)
			xcoord, ycoord := CalcRange(chunk, x)
			fmt.Println(xcoord)
			fmt.Println(ycoord)
		}
	}
}

// Day 3, Part 2 of the AoC2023 challenge
func Part2() {
	fmt.Println("Day 3, Part 2: Gear Ratios")
}
