package main

import (
	"aoc2023/src/day_1"
	"flag"
	"fmt"
)

func main() {
	testPtr := flag.Bool("test", false, "enable test mode")
	dayPtr := flag.Int("day", 1, "which day to run")
	partPtr := flag.Int("part", 1, "which part to run")

	flag.Parse()

	fmt.Println("Advent of Code 2023!")
	fmt.Println("  test:", *testPtr)
	fmt.Println("  day:", *dayPtr)
	fmt.Println("  part:", *partPtr)

	if *dayPtr == 1 {
		if *partPtr == 1 {
			day_1.Part1()
		} else {
			day_1.Part2()
		}

	}
}
