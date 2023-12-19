package day4

import (
	"aoc2023/src/base"
	"fmt"
	"math"

	"github.com/juliangruber/go-intersect"
)

type card struct {
	winners []int
	numbers []int
}

// Uses exponention base 2 function to determine the number of points the winners get
func calcPoints(winners int) int {
	if winners == 0 {
		return 0
	}
	return int(math.Exp2(float64(winners - 1)))
}

func calcWinners(card card) int {
	return len(intersect.Simple(card.winners, card.numbers))
}

// Converts a string into a card type
// Example card:
// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
func parseLine(line string) card {
	var card card

	// disregard everything prior to :
	// nums := strings.Split(line, ":")[1]
	// w := strings.Split(nums, "|")[0]
	// n := strings.Split(nums, "|")[1]

	return card
}

func Part1() {
	var linecount, sum int

	file := base.GetDayDataFile(4, 1)

	lines, err := base.ReadLines(file)
	if err != nil {
		panic(err)
	}

	for n, line := range lines {
		card := parseLine(line)
		fmt.Println(n, "card:", card)
	}

	fmt.Println("Day 4, Part 1: Scratchcards")
	fmt.Println("Lines processed", linecount)
	fmt.Println("Sum of adjacent chunks", sum)
}

func Part2() {
	var linecount, sum int
	fmt.Println("Day 4, Part 2: Scratchcards")
	fmt.Println("Lines processed", linecount)
	fmt.Println("Sum of adjacent chunks", sum)
}
