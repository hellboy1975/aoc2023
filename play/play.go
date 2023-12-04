package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func getNumbers(x string) string {
	reg := regexp.MustCompile(`\D`)
	return reg.ReplaceAllString(x, "")
}

func main() {
	s := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"

	re := regexp.MustCompile(`\d+:`)
	fmt.Printf("%q\n", re.FindString(s))

	rg := strings.Split(s, ":")[0]
	fmt.Printf("%q\n", rg)
	fmt.Println(getNumbers(rg))

	i, _ := strconv.Atoi(rg)

	fmt.Println(i)

}
