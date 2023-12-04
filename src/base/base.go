package base

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

const testDataFile = "data/day_%d_%d.txt"

// gets the data file for the requested day and part
func GetDayDataFile(day, part int) string {
	file := fmt.Sprintf(testDataFile, day, part)
	fmt.Println("  Data file: " + file)

	return file
}

// Read a whole file into the memory and store it as array of lines
func ReadLines(path string) (lines []string, err error) {
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
