package day3

import (
	"aoc2023/src/base"
	"errors"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"unicode"
)

const CharNum = 1
const CharSymbol = 2
const CharBlank = 3
const CharGear = 4

const DataExtent = 139 // 140 columns/rows in the data provided

const NumRegex = "\\d+"

type coord struct {
	x int
	y int
}

var symbols [][]bool
var gears [][]bool

// returns 1 for number, 2 for symbol and 3 for .
func CharType(c rune) int {
	if c == '*' {
		return CharGear
	} else if c == '.' {
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
			// Gears and other symbols are different
			celltype := CharType(rune(cell))
			if celltype == CharSymbol || celltype == CharGear {
				sline[x] = true
			} else {
				sline[x] = false
			}
		}
		symbols[y] = sline
	}
	return symbols
}

// Creates a 2D array of boolean values indicating if a x,y coordinate is a symbol
func Gears(lines []string) [][]bool {
	symbols := make([][]bool, 140)

	for y, line := range lines {
		sline := make([]bool, 140)
		for x, cell := range line {
			if CharType(rune(cell)) == CharGear {
				sline[x] = true
			} else {
				sline[x] = false
			}
		}
		symbols[y] = sline
	}
	return symbols
}

// returns the integer value of the chunk located on the passed cell
func GetChunkValue(cell int, chunks [][]int, line string) (int, error) {
	for _, chunk := range chunks {
		if cell >= chunk[0] && cell < chunk[1] {
			// get the whole number
			r, _ := strconv.Atoi(line[chunk[0]:chunk[1]])
			return r, nil
		}
	}
	return -1, errors.New("no chunk match")
}

// iterates over the file, building a 2D array in which each element contains either 0 or the number of the chunk
func Chunks(lines []string) (data [][]int) {
	var linecount int

	for l, line := range lines {
		fmt.Println(line)
		linecount++
		chunks := LineNumberChunks(line)

		for x, cell := range line {
			// fmt.Println(l, x, cell)
			if unicode.IsDigit(cell) {
				// if it's a number, then find the matching chunk
				c, err := GetChunkValue(x, chunks, line)
				if err != nil {
					panic(err)
				}
				data[l][x] = c
			} else {
				data[l][x] = -1
			}
		}
	}
	return
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
	// iterate each of the rows between the x & y y coordinates
	for i := x.y; i <= y.y; i++ {

		s := symbols[i][x.x : y.x+1]
		if slices.Contains(s, true) {
			return true
		}
	}

	return false
}

// Visualises the symbols
func PrintGrid(grid [][]bool) {
	for _, line := range grid {
		fmt.Print("|")
		for _, s := range line {
			if s {
				fmt.Print("*")
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

	PrintGrid(symbols)

	var sum, count, linecount int

	for x, line := range lines {
		linecount++
		chunks := LineNumberChunks(line)

		for _, chunk := range chunks {
			count++
			xcoord, ycoord := CalcRange(chunk, x)
			if IsChunkNextToSymbol(xcoord, ycoord) {
				add, _ := strconv.Atoi(string(line[chunk[0]:chunk[1]]))
				sum += add
			}

		}
	}
	fmt.Println("Lines processed", linecount)
	fmt.Println("Sum of adjacent chunks", sum)
}

// Day 3, Part 2 of the AoC2023 challenge
func Part2() {
	fmt.Println("Day 3, Part 2: Gear Ratios")

	file := base.GetDayDataFile(3, 1)

	lines, err := base.ReadLines(file)
	if err != nil {
		panic(err)
	}

	gears = Gears(lines)

	PrintGrid(gears)

	chunks := Chunks(lines)

	fmt.Println("chunks[0]", chunks[0])
}
