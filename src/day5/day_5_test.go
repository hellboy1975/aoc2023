package day5

import (
	"reflect"
	"testing"
)

func init() {
}

func TestParseSeeds(t *testing.T) {

	cases := []struct {
		Description string
		Raw         string
		Want        []int
	}{
		{"Example seeds", "seeds: 79 14 55 13", []int{79, 14, 55, 13}},
		{"My seeds",
			"seeds: 3943078016 158366385 481035699 103909769 3553279107 15651230 3322093486 189601966 2957349913 359478652 924423181 691197498 2578953067 27362630 124747783 108079254 1992340665 437203822 2681092979 110901631",
			[]int{3943078016, 158366385, 481035699, 103909769, 3553279107, 15651230, 3322093486, 189601966, 2957349913, 359478652, 924423181, 691197498, 2578953067, 27362630, 124747783, 108079254, 1992340665, 437203822, 2681092979, 110901631}},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ParseSeeds(test.Raw)
			if reflect.DeepEqual(got, test.Want) != true {
				t.Errorf("got %d, want %d", got, test.Want)
			}
		})
	}
}
