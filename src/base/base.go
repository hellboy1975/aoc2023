package base

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
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

// Read a whole file into the memory and store it as array of lines
func readLines(path string) (lines []string, err error) {
	var (
		file   *os.File
		part   []byte
		prefix bool
	)
	if file, err = os.Open(path); err != nil {
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buffer := bytes.NewBuffer(make([]byte, 0))
	for {
		if part, prefix, err = reader.ReadLine(); err != nil {
			break
		}
		buffer.Write(part)
		if !prefix {
			lines = append(lines, buffer.String())
			buffer.Reset()
		}
	}
	if err == io.EOF {
		err = nil
	}
	return
}
