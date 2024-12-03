package main

import (
	"math"
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
	// when you're ready to do part 2, remove this "not implemented" bloc
	if part2 {
		return "not implemented"
	}
	// solve part 1 here

	at_least_one := 1
	at_most_three := 3

	data := strings.Split(input, "\n")

	for _, line := range data {
		columns := strings.SplitN(line, " ", len(line))
		preveous := 0
		for _, numbers := range columns {
			i, err := strconv.Atoi(numbers)
			if err != nil {
				// ... handle error
				panic(err)
			}
			if preveous != 0 {
				diff := preveous - i
				print(preveous)
				print("-")
				print(i)
				print("=")
				print(diff)
				if math.Abs(float64(diff)) > float64(at_most_three) {
					println("_Unsafe_")
					break
				} else if math.Abs(float64(diff)) < float64(at_least_one) {
					println("_Unsafe_")
					break
				} else {
					println("_Safe_")
				}
			}
			preveous = i

		}
		//
	}
	return 42
}
