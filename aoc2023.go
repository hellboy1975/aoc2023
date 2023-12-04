package main

import (
	"aoc2023/src/day_1"
	"flag"
	"fmt"
)

func main() {
	dayPtr := flag.Int("day", 1, "which day to run")
	partPtr := flag.Int("part", 1, "which part to run")

	flag.Parse()

	fmt.Println("Advent of Code 2023!")

	if *dayPtr == 1 {
		if *partPtr == 1 {
			day_1.Part1()
		} else {
			day_1.Part2() // pass an argument to reduce code reuse
		}

	}
}
