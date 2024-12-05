package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	startTime := time.Now()

	input, err := os.ReadFile("04.1/exemple.txt")
	if err != nil {
		fmt.Println("Error Opening File")
		os.Exit(1)
	}

	var nXmas int

	// var lineStart int
	lineSize := 300
	for i, j := 0, 0; i < len(input); i, i = i+1, j+1 {
		directions := [][]int{{-1, 0}, {-1, -1}, {0, -1}, {1, 1}}

		if input[i] == 10 {
			lineSize = j + 1
			j = -1
			fmt.Print(lineSize)
			continue
		}

		for _, dir := range directions {
			xMin := j + 3*dir[0]
			yMin := i + 3*dir[1]*lineSize
			if xMin >= 0 && xMin <= lineSize && yMin >= 0 {
				if (input[i+0*dir[0]+0*dir[1]*lineSize] == 'X' &&
					input[i+1*dir[0]+1*dir[1]*lineSize] == 'M' &&
					input[i+2*dir[0]+2*dir[1]*lineSize] == 'A' &&
					input[i+3*dir[0]+3*dir[1]*lineSize] == 'S') ||
					(input[i+0*dir[0]+0*dir[1]*lineSize] == 'S' &&
						input[i+1*dir[0]+1*dir[1]*lineSize] == 'A' &&
						input[i+2*dir[0]+2*dir[1]*lineSize] == 'M' &&
						input[i+3*dir[0]+3*dir[1]*lineSize] == 'X') {
					nXmas++
				}
			}
		}
	}

	fmt.Println(fmt.Sprint("The total number of XMAS is: ", nXmas))

	fmt.Println(fmt.Sprint("Script took: ", time.Now().Sub(startTime)))
}
