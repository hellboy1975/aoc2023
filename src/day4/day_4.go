package day4

import (
	"aoc2023/src/base"
	"fmt"
	"math"
	"strings"

	"github.com/juliangruber/go-intersect"
)

type card struct {
	id        int
	winners   []int
	numbers   []int
	wins      int
	collected int
}

var cards []card

var loopcount int

// Uses exponention base 2 function to determine the number of points the winners get
func calcPoints(winners int) int {
	if winners == 0 {
		return 0
	}
	return int(math.Exp2(float64(winners - 1)))
}

// compares the number and winner arrays and returns the number of matching values
func calcWinners(card card) int {
	return len(intersect.Simple(card.winners, card.numbers))
}

// takes a look at the passed card, and based on how many are collected (always at least one) it then adds extra cards
func calcCardCollection(card card) {
	cardcount := len(cards)

	// you always have at least the original card
	cards[card.id-1].collected += 1

	fmt.Println("Processing:", card.id, "collected", cards[card.id-1].collected)

	// if there are no winners no further processing is required
	if card.wins > 0 {
		// for each card collected
		for c := 0; c < cards[card.id-1].collected; c++ {
			loopcount++

			// process cards for each winning card
			for i := card.id; i < card.id+card.wins; i++ {
				loopcount++
				if i >= cardcount {
					break
				}

				cards[i].collected += 1
			}
		}
	}

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
		loopcount++
		card := parseLine(line)
		card.wins = calcWinners(card)
		card.id = linecount

		cards = append(cards, card)
	}

	// second pass will count how many cards are won
	for _, card := range cards {
		loopcount++
		calcCardCollection(card)
	}

	// third pass adds them altogether
	for _, card := range cards {
		loopcount++
		sum += card.collected
	}

	fmt.Println("Day 4, Part 2: Scratchcards")
	fmt.Println("Lines processed", linecount)
	fmt.Println("Loops ran", loopcount)
	fmt.Println("Number of cards won:", sum)
}
