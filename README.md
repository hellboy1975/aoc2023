# Advent of Code 2023

A bunch of code I've create to compete in the 2023 Advent of Code challenge.

My approach for this year:

* Learn how to write in Go
* Use TDD

## Usage

Clone this repo to your file system directory of choice, and make sure you have Go 1.21.4 installed

Build the project with:

```
go build
```

Run a specific day with a command like:

```
./aoc2023 -day=1 -part=1
```

For limited help:

```
./aoc2023 -h
Usage of ./aoc2023:
  -day int
        which day to run (default 1)
  -part int
        which part to run (default 1)
```

Run all of the tests with:

```
go test ./...
```

## Day notes

### Day 1

Part 1 was relatively straight forward.

Part 2 was significantly harder.  Initial approach of just a string replacement of the number words failed, as some of the data contained overlapping numbers.
For example: `1twoxyzoneeight`

I came up with a second approach that used a fancy regex using a look ahead:

```
(?=(one|two|three|four|five|six|seven|eight|nine|1|2|3|4|5|6|7|8|9|0))
```

I'm pretty confident this would have worked, however the various regex libraries available for Go don't support this syntax.  As I'm keen to move on to Day 2,
I have decided refactoring this (again) isn't worth the effort)

## References

Below is a list of sites/resources that have been useful in these challenges:

* https://regex101.com/
* https://gobyexample.com
* https://quii.gitbook.io/learn-go-with-tests/
