package day2

import (
	"testing"
)

func TestGetGameId(t *testing.T) {

	cases := []struct {
		Description string
		Raw         string
		Want        int
	}{
		{"'Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green' extracts ID of 1", "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", 1},
		{"'Game 612: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green' extracts ID of 1", "Game 612: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", 612},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := GetGameId(test.Raw)
			if got != test.Want {
				t.Errorf("got %d, want %d", got, test.Want)
			}
		})
	}
}

func TestGetColourMax(t *testing.T) {

	cases := []struct {
		Description string
		Raw         string
		Blue        int
		Red         int
		Green       int
	}{
		{"Game 1", "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", 6, 4, 2},
		{"Game 2", "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", 4, 1, 3},
		{"Game 3", "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", 6, 20, 13},
		{"Game 4", "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", 15, 14, 3},
		{"Game 5", "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", 2, 6, 3},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			blue, red, green := GetColourMax(test.Raw)

			if blue != test.Blue {
				t.Errorf("Blue: got %d, want %d", blue, test.Blue)
			}
			if red != test.Red {
				t.Errorf("Red: got %d, want %d", red, test.Red)
			}
			if green != test.Green {
				t.Errorf("Green: got %d, want %d", green, test.Green)
			}
		})
	}
}

func TestIsGamePossible(t *testing.T) {

	cases := []struct {
		Description string
		Raw         string
		Want        bool
	}{
		{"Game 1 is possible", "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", true},
		{"Game 2 is possible", "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", true},
		{"Game 3 is impossible", "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", false},
		{"Game 4 is impossible", "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", false},
		{"Game 5 is possible", "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", true},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := IsGamePossible(test.Raw)
			if got != test.Want {
				t.Errorf("got %t, want %t", got, test.Want)
			}
		})
	}
}
