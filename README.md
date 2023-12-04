# aoc2023
Advent of Code 2023

## Day notes

### Day 1

Part 1 was relatively straight forward.

Part 2 was significantly harder.  Initial approach of just a string replacement of the number words failed, as some of the data contained overlapping numbers.
For example: `1twoxyzoneeight`

I came up with a second approach that used a fancy regex using a look ahead:

```(?=(one|two|three|four|five|six|seven|eight|nine|1|2|3|4|5|6|7|8|9|0))```

I'm pretty confident this would have worked, however the various regex libraries available for Go don't support this syntax.  As I'm keen to move on to Day 2,
I have decided refactoring this (again) isn't worth the effort)

## References

Below is a list of sites/resources that have been useful in these challenges:

* https://regex101.com/
* https://gobyexample.com
* https://quii.gitbook.io/learn-go-with-tests/
