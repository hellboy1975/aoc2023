package main

import (
	"aoc2023/src/day1"
	"aoc2023/src/day2"
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
			day1.Part1()
		} else {
			day1.Part2() // pass an argument to reduce code reuse
		}
	} else if *dayPtr == 2 {
		if *partPtr == 1 {
			day2.Part1()
		} else {
			day2.Part2() // pass an argument to reduce code reuse
		}
	}
}
