package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	startTime := time.Now()

	input, err := os.ReadFile("06.1/input.txt")
	if err != nil {
		fmt.Println("Error Opening File")
		os.Exit(1)
	}

	var mapping [][]byte
	var gardPosition [2]int
	rotations := [4][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	for y := 0; y < len(input); y++ {
		for x := 0; y < len(input); y, x = y+1, x+1 {
			if input[y] == '^' {
				gardPosition = [2]int{x, (y / len(mapping[0]))}
			}

			if input[y] == '\n' {
				mapping = append(mapping, input[y-x:y])
				break
			}
		}
	}

	gardOut := false
	currentPosition := gardPosition
	currentRotation := 0
	gardPath := make(map[int]map[int]bool)
	var nPositions int
	for i := 0; i < len(input); i++ {
		gardPath[i] = make(map[int]bool)
	}

	for !gardOut {
		if currentPosition[0] > len(mapping[0])-1 || currentPosition[1] > len(mapping)-1 || currentPosition[0] < 0 || currentPosition[1] < 0 {
			gardOut = true
			break
		}

		if mapping[currentPosition[1]][currentPosition[0]] == '#' {
			currentPosition[0] -= 1 * rotations[currentRotation%4][0]
			currentPosition[1] -= 1 * rotations[currentRotation%4][1]
			currentRotation++
			continue
		}

		if !gardPath[currentPosition[0]][currentPosition[1]] {
			gardPath[currentPosition[0]][currentPosition[1]] = true
			nPositions++
		}

		currentPosition[0] += 1 * rotations[currentRotation%4][0]
		currentPosition[1] += 1 * rotations[currentRotation%4][1]
	}

	fmt.Println("The number of positions is: ", nPositions)

	fmt.Println(fmt.Sprint("Script took: ", time.Now().Sub(startTime)))
}
