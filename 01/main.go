package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	fmt.Print("Reading input file...")

	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error Opening File")
		os.Exit(1)
	}
	fmt.Print("OK\n")

	defer f.Close()

	scanner := bufio.NewScanner(f)

	fmt.Print("Parsing File data...")
	var leftValues, rightValues []int
	for scanner.Scan() {
		var leftValue, rightValue int
		_, err = fmt.Sscanf(scanner.Text(), "%d   %d", &leftValue, &rightValue)
		if err != nil {
			fmt.Print("Error parsing line values\n")
			os.Exit(1)
		}
		leftValues = append(leftValues, leftValue)
		rightValues = append(rightValues, rightValue)
	}
	fmt.Print("OK\n")

	fmt.Print("Sorting Data...")
	sort.Ints(leftValues)
	sort.Ints(rightValues)
	fmt.Print("OK\n")

	fmt.Print("Calculating final sum...")
	totalLengh := 0
	for i := 0; i < len(leftValues); i++ {
		totalLengh += int(math.Abs(float64(leftValues[i] - rightValues[i])))
	}
	fmt.Print("OK\n")

	fmt.Println(fmt.Sprint("The total lenght is: ", totalLengh))
}
