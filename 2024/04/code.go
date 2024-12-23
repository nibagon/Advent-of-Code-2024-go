package main

import (
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
func find_index(i int, j int, row_len int) int {
	index := (i * row_len) + (i * 1) + j
	return index
}
func check(input string, word string, i int, j int, n int, m int, direction int) bool {
	substring := ""
	for ((0 <= i) && (i < n)) && ((0 <= j) && (j < m)) && (len(substring) <= len(word)) {
		substring += string(input[find_index(i, j, n)])
		if substring == word {
			return true
		}
		switch direction {
		case 0:
			j += 1
		case 1:
			j += 1
			i += 1
		case 2:
			i += 1
		case 3:
			i += 1
			j -= 1
		case 4:
			j -= 1
		case 5:
			i -= 1
			j -= 1
		case 6:
			i -= 1
		case 7:
			i -= 1
			j += 1
		}
		if i < 0 || j < 0 || i == n || j == m {
			return false
		}
	}
	return false
}
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}
	// solve part 1 here
	word := "XMAS"
	// substring:=""

	// for i := 0; i < len(input); i++ {
	// 	println(string(input[i]))

	// }

	row := strings.Split(input, "\n")
	directions := 8
	wordscount := 0
	xcount := 0
	// word_pointer ;= 0
	for i, line := range row {
		column := strings.Split(line, "")
		for j, letter := range column {
			if letter == "X" {
				xcount++
				for direction := 0; direction < directions; direction++ {
					if check(input, word, i, j, len(row), len(column), direction) {
						wordscount++
					}
				}
			}
		}
	}
	println()
	println("found ", wordscount, " XMAS and ", xcount, " X")

	return 42
}
