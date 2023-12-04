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
