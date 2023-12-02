package base

import (
	"fmt"
)

const testDataFile = "day_%d_%d%s.txt"

func GetTestDataFile(test bool, day, part int) string {
	var replace string

	if test {
		replace = "_test"
	}

	testDataFile := fmt.Sprintf(testDataFile, day, part, replace)
	fmt.Println("  Test data file: " + testDataFile)

	return testDataFile
}
