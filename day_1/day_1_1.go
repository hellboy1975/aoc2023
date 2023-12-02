package main

import (
	"fmt"
	"regexp"
)

func testData() [][]string {
	return [][]string{
		{"1abc2"},
		{"pqr3stu8vwx"},
		{"a1b2c3d4e5f"},
		{"treb7uchet"},
	}
}

func getNumbers(string x) string {
	reg := regexp.MustCompile("[0-9]")
	return reg.ReplaceAllString(x, "")
}

func main() {
	fmt.Println("Day 1: Trebuchet?!")

	data := testData()

}
