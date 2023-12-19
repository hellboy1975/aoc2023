package day4

import "testing"

func init() {
}

func TestCalcPoints(t *testing.T) {

	cases := []struct {
		Description string
		Raw         int
		Want        int
	}{
		{"0 winners is 0 points", 0, 0},
		{"1 winners is 1 points", 1, 1},
		{"2 winners is 2 points", 2, 2},
		{"3 winners is 6 points", 3, 4},
		{"4 winners is 12 points", 4, 8},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := calcPoints(test.Raw)
			if got != test.Want {
				t.Errorf("got %d, want %d", got, test.Want)
			}
		})
	}
}

func TestCalcWinners(t *testing.T) {

	cases := []struct {
		Description string
		card        card
		Want        int
	}{
		{"3 items resolves to 2 winners", card{winners: []int{1, 2, 4}, numbers: []int{2, 4}}, 2},
		{"3 items (reversed) resolves to 2 winners", card{winners: []int{2, 4}, numbers: []int{1, 2, 4}}, 2},
		{"Card 3 example is 2 winners", card{winners: []int{1, 21, 53, 59, 44}, numbers: []int{69, 82, 63, 72, 16, 21, 14, 1}}, 2},
		{"Card 6 example is 0 winners", card{winners: []int{31, 18, 13, 56, 72}, numbers: []int{74, 77, 10, 23, 35, 67, 36, 11}}, 0},
		{"5 dimension matching array is 5 winners", card{winners: []int{1, 2, 3, 4, 5}, numbers: []int{1, 2, 3, 4, 5}}, 5},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := calcWinners(test.card)
			if got != test.Want {
				t.Errorf("got %d, want %d", got, test.Want)
			}
		})
	}
}
