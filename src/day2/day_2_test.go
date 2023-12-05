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

func TestGetColourCount(t *testing.T) {

	wantblue := 9
	wantred := 5
	wantgreen := 4

	blue, red, green := GetColourCount("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")

	if blue != wantblue {
		t.Errorf("blue expected '%d' but got '%d'", wantblue, blue)
	}
	if red != wantred {
		t.Errorf("red expected '%d' but got '%d'", wantred, red)
	}
	if green != wantgreen {
		t.Errorf("green expected '%d' but got '%d'", wantgreen, green)
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
