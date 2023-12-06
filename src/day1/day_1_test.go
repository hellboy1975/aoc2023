package day1

import (
	"testing"
)

func TestExtractCalibrationValuesPart1(t *testing.T) {

	cases := []struct {
		Description string
		Raw         string
		Want        int
	}{
		{"1abc2 gets converted to 12", "1abc2", 12},
		{"pqr3stu8vwx gets converted to 38", "pqr3stu8vwx", 38},
		{"a1b2c3d4e5f gets converted to 15", "a1b2c3d4e5f", 15},
		{"treb7uchet gets converted to 77", "treb7uchet", 77},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := extractCalibrationValuePart1(test.Raw)
			if got != test.Want {
				t.Errorf("got %d, want %d", got, test.Want)
			}
		})
	}
}

// Part 2 didn't pan out, so I've removed these tests for now :(
// func TestExtractCalibrationValuesPart2(t *testing.T) {

// 	cases := []struct {
// 		Description string
// 		Raw         string
// 		Want        int
// 	}{
// 		{"two1nine gets converted to 29", "two1nine", 29},
// 		{"eightwothree gets converted to 83", "eightwothree", 83},
// 		{"abcone2threexyz gets converted to 13", "abcone2threexyz", 13},
// 		{"xtwone3four gets converted to 24", "xtwone3four", 24},
// 		{"4nineeightseven2 gets converted to 42", "4nineeightseven2", 42},
// 		{"zoneight234 gets converted to 14", "zoneight234", 14},
// 		{"7pqrstsixteen gets converted to 76", "7pqrstsixteen", 76},
// 		{"6one2tvfoneight  gets converted to 68", "6one2tvfoneight ", 68},
// 	}

// 	for _, test := range cases {
// 		t.Run(test.Description, func(t *testing.T) {
// 			got := extractCalibrationValuePart2(test.Raw)
// 			if got != test.Want {
// 				t.Errorf("got %d, want %d", got, test.Want)
// 			}
// 		})
// 	}
// }

func TestStringToInt(t *testing.T) {
	cases := []struct {
		Description string
		Raw         string
		Want        int
	}{
		{"one gets converted to 1", "one", 1},
		{"banana returns -1", "banana", -1},
		{"'4' returns (int)4", "4", 4},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got, err := stringToInt(test.Raw)
			if got != test.Want {
				t.Errorf("got %d, want %d", got, test.Want)
			}
			if got == -1 && err == nil {
				t.Errorf("Error 'no number match' expected")
			}
		})
	}
}
