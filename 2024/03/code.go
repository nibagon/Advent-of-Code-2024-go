package main

import (
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}
	// solve part 1 here
	data := strings.Split(input, "mul")
	print("I found muls:")
	println(len(data))
	sum := 0
	for _, line := range data {
		counter := 0
		first_number := 0
		for _, element := range strings.Split(line, "") {
			i, err := strconv.Atoi(element)
			if err != nil {
				// ... handle error
				continue
			}
			if counter == 0 {
				first_number = i
				counter++
			} else if counter == 1 {
				mul := i * first_number
				print("mul ")
				println(i, first_number)
				sum += mul
				counter = 0
			} else {
				counter = 0
			}
		}
	}
	print("sum =")
	println(sum)

	return 42
}
