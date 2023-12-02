package day_1

import (
	"testing"
)

func TestExtractCalibrationValues(t *testing.T) {

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
			got := extractCalibrationValue(test.Raw)
			if got != test.Want {
				t.Errorf("got %d, want %d", got, test.Want)
			}
		})
	}
}
