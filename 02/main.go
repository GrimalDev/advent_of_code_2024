package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Reading input file...")

	f, err := os.Open("01/input.txt")
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

	rightMap := make(map[int]int)

	for _, id := range rightValues {
		rightMap[id]++
	}

	similarityScore := 0
	for _, id := range leftValues {
		similarityScore += id * rightMap[id]
	}

	fmt.Println(fmt.Sprint("Similarity score: ", similarityScore))
}
