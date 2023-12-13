package day3

import (
	"aoc2023/src/base"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"unicode"
)

const CharNum = 1
const CharSymbol = 2
const CharBlank = 3

const DataExtent = 139 // 140 columns/rows in the data provided

// const NumRegex = "([0-9])\\w+"
const NumRegex = "\\d+"

type coord struct {
	x int
	y int
}

var symbols [][]bool

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
	symbols := make([][]bool, 140)

	for y, line := range lines {
		sline := make([]bool, 140)
		for x, cell := range line {
			if CharType(rune(cell)) == CharSymbol {
				sline[x] = true
			} else {
				sline[x] = false
			}
		}
		symbols[y] = sline
	}
	return symbols
}

func Ycoord(chunk []int, line int) (ycoord coord) {
	// ycoord will contain the bottom-right corner of the range "box"
	if chunk[1] >= DataExtent {
		ycoord.x = DataExtent
	} else {
		ycoord.x = chunk[1]
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

// checks to see if the chunk has a symbol with it's box as defined by x & y
func IsChunkNextToSymbol(x coord, y coord) bool {
	fmt.Println("IsChunkNextToSymbol", x.x, x.y, y.x, y.y)
	fmt.Println("symbols", len(symbols))
	// iterate each of the rows between the x & y y coordinates
	for i := x.y; i <= y.y; i++ {
		fmt.Println("line:", i)
		s := symbols[i][x.x:y.x]
		if slices.Contains(s, true) {
			return true
		}
	}

	return false
}

// Visualises the symbols
func PrintSymbols() {
	fmt.Println("Symbols grid:")
	for _, line := range symbols {
		fmt.Print("|")
		for pos, s := range line {
			if s {
				fmt.Print("*")
				fmt.Print(pos)
			} else {
				fmt.Print(".")
			}

		}
		fmt.Println("|")
	}
}

// Day 3, Part 1 of the AoC2023 challenge
func Part1() {
	fmt.Println("Day 3, Part 1: Gear Ratios")

	file := base.GetDayDataFile(3, 1)

	lines, err := base.ReadLines(file)
	if err != nil {
		panic(err)
	}

	symbols = Symbols(lines)

	PrintSymbols()

	var sum, count, linecount int

	for x, line := range lines {
		linecount++
		chunks := LineNumberChunks(line)

		for _, chunk := range chunks {
			count++
			// fmt.Println("line", l, chunk)
			xcoord, ycoord := CalcRange(chunk, x)
			// fmt.Println(xcoord)
			// fmt.Println(ycoord)
			if IsChunkNextToSymbol(xcoord, ycoord) {

				if x == 139 {
					fmt.Println(xcoord, ycoord)
					fmt.Println(count, ": Chunk", chunk, "is adjacent.  Line:", x)
				}
				add, _ := strconv.Atoi(string(line[chunk[0]:chunk[1]]))
				sum += add
			} else {
				if x == 139 {
					fmt.Println(xcoord, ycoord)
					fmt.Println(count, ": Chunk", chunk, "is not adjacent.  Line:", x)
				}
			}

		}
	}
	fmt.Println("Lines processed", linecount)
	fmt.Println("Sum of adjacent chunks", sum)
}

// Day 3, Part 2 of the AoC2023 challenge
func Part2() {
	fmt.Println("Day 3, Part 2: Gear Ratios")
}
