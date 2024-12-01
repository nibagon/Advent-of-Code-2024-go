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

func sortArray(arr [1000]int) [1000]int {
	for i := 0; i <= len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}
	// solve part 1 here
	// first open the file
	//Section 1
	sum := 0
	data := strings.Split(input, "\n")

	// numberOfIds := len(data)
	// println(numberOfIds)
	leftArray := [1000]int{}
	rigthArray := [1000]int{}
	for _, line := range data {
		columns := strings.SplitN(line, "   ", 2)
		il, err := strconv.Atoi(columns[0])
		if err != nil {
			// ... handle error
			panic(err)
		}
		ir, err := strconv.Atoi(columns[1])
		if err != nil {
			// ... handle error
			panic(err)
		}
		leftArray[sum] = il
		rigthArray[sum] = ir
		sum++
		// println(pair[1])
	}
	resultleft := sortArray(leftArray)
	resultright := sortArray(rigthArray)
	// slices.Sort(leftArray)
	// slices.Sort(rigthArray)
	sum = 0
	diff := 0
	for i := 0; i < len(leftArray); i++ {

		diff = (resultleft[i] - resultright[i])
		println(resultleft[i], "-", resultright[i], "=", diff)
		intDiff := float64(diff)
		intDiff = math.Abs(intDiff)
		sum = sum + int(intDiff)

	}
	println("Solution =", sum)

	return sum
}
