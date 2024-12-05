package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	startTime := time.Now()

	input, err := os.ReadFile("04.1/input.txt")
	if err != nil {
		fmt.Println("Error Opening File")
		os.Exit(1)
	}

	var nXmas int
	lines := [][]byte{}
	currentLine := []byte{}

	// Split input into lines
	for _, b := range input {
		if b == '\n' {
			lines = append(lines, currentLine)
			currentLine = []byte{}
		} else {
			currentLine = append(currentLine, b)
		}
	}
	if len(currentLine) > 0 {
		lines = append(lines, currentLine)
	}

	directions := [][]int{{-1, 0}, {-1, -1}, {0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}}

	for y, line := range lines {
		for x := range line {
			for _, dir := range directions {
				if checkXMAS(lines, x, y, dir[0], dir[1]) {
					nXmas++
				}
			}
		}
	}

	fmt.Printf("The total number of XMAS is: %d\n", nXmas)
	fmt.Printf("Script took: %v\n", time.Since(startTime))
}

func checkXMAS(lines [][]byte, x, y, dx, dy int) bool {
	word := "XMAS"
	for i := 0; i < 4; i++ {
		nx, ny := x+i*dx, y+i*dy
		if ny < 0 || ny >= len(lines) || nx < -1 || nx >= len(lines[ny]) {
			return false
		}
		if lines[ny][nx] != word[i] {
			return false
		}
	}
	return true
}
