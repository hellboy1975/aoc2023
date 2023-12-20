package day4

import (
	"aoc2023/src/base"
	"fmt"
	"math"
	"strings"

	"github.com/juliangruber/go-intersect"
)

type card struct {
	id      int
	winners []int
	numbers []int
	wins    int
}

var cards []card

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

// func calcCardWins(card card) int {
// 	var wins int

// 	if card.id == len(cards) {
// 		return 0
// 	}

// 	fmt.Println("card process:", card.id, card.wins, "len", len(cards))

// 	wins = card.wins

// 	fmt.Print("For Card ", card.id, " loop ")
// 	for i := card.id; i <= card.wins+card.id-1; i++ {
// 		fmt.Print(" [", i, "]")
// 		wins += calcCardWins(cards[i])
// 	}
// 	fmt.Println()
// 	fmt.Println("card wins:", card.id, wins)
// 	return wins
// }

func calcCardWins(card card) int {
	wins := card.wins

	fmt.Println("Parent card", card.id)
	for i := card.id; i < card.id+card.wins; i++ {
		fmt.Println("Iterate card", i)
	}
	return wins
}

// Converts a string into a card type
// Example card:
// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
func parseLine(line string) card {
	var card card

	// disregard everything prior to :
	nums := strings.Split(line, ":")[1]

	card.winners = base.StringToIntArray(strings.Split(nums, "|")[0])
	card.numbers = base.StringToIntArray(strings.Split(nums, "|")[1])

	return card
}

func Part1() {
	var linecount, sum int

	file := base.GetDayDataFile(4, 1)

	lines, err := base.ReadLines(file)
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		linecount++
		card := parseLine(line)

		points := calcPoints(calcWinners(card))

		sum += points
	}

	fmt.Println("Day 4, Part 1: Scratchcards")
	fmt.Println("Lines processed", linecount)
	fmt.Println("Sum of winning points", sum)
}

func Part2() {
	var linecount, sum int

	file := base.GetDayDataFile(4, 1)

	lines, err := base.ReadLines(file)
	if err != nil {
		panic(err)
	}

	// first pass will setup the cards array
	for _, line := range lines {
		linecount++
		card := parseLine(line)
		card.wins = calcWinners(card)
		card.id = linecount

		cards = append(cards, card)
	}

	for i, card := range cards {
		fmt.Println("card:", card.id, "----------------------------")
		won := calcCardWins(card)
		fmt.Println("card:", card.id, card.wins, "won:", won)

		if i > 10 {
			break
		}
		sum += won
	}

	// second pass with traverse the cards, and work out how many winners each one collects

	fmt.Println("Day 4, Part 2: Scratchcards")
	fmt.Println("Lines processed", linecount)
	fmt.Println("Number of cards won:", sum)
}
