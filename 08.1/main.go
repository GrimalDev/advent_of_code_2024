package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func getAntinode(A [2]int, B [2]int, iteration int) [2]int {
	direction := [2]int{B[0] - A[0], B[1] - A[1]}
	return [2]int{B[0] + direction[0]*iteration, B[1] + direction[1]*iteration}
}

func main() {
	startTime := time.Now()

	input, err := os.Open("08.1/input.txt")
	if err != nil {
		fmt.Println("Error Opening File")
		os.Exit(1)
	}

	defer input.Close()

	scanner := bufio.NewScanner(input)

	antenas := make(map[rune][][2]int)
	antinodes := make(map[[2]int]bool)

	var maxX int
	var maxY int

	for scanner.Scan() {
		line := scanner.Text()
		for x, char := range line {
			if char != '.' {
				antenas[char] = append(antenas[char], [2]int{x, maxY})
			}
		}
		if maxX == 0 {
			maxX = len(scanner.Text())
		}
		maxY++
	}

	maxY--
	maxX--

	for _, positions := range antenas {
		for _, position1 := range positions {
			for _, position2 := range positions {
				if position1 == position2 {
					antinodes[position2] = true
					continue
				}
				for iteration := 1; iteration > 0; iteration++ {
					antinode := getAntinode(position1, position2, iteration)
					if antinode[0] < 0 || antinode[0] > maxX || antinode[1] < 0 || antinode[1] > maxY {
						break
					}
					antinodes[antinode] = true
				}
			}
		}
	}

	fmt.Println("The number of antinodes is: ", len(antinodes))

	fmt.Println(fmt.Sprint("Script took: ", time.Now().Sub(startTime)))
}
