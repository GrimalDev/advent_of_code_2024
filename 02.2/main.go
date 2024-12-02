package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func reverse(values []string) []string {
	var rev []string
	for _, n := range values {
		rev = append([]string{n}, rev...)
	}
	return rev
}

func isSafe(reportRAW []string) bool {
	var report []int
	for _, levelRAW := range reportRAW {
		level, _ := strconv.Atoi(levelRAW)
		report = append(report, level)
	}

	var direction int
	secondChance := false
	tried := false
	for i := 1; i < len(report); i++ {
		var newDirection int
		if secondChance && !tried {
			newDirection = report[i] - report[i-2]
			tried = true
		} else {
			newDirection = report[i] - report[i-1]
		}
		delta := math.Abs(float64(newDirection))

		if delta > 3 || delta < 1 {
			if secondChance {
				return false
			}
			secondChance = true
			continue
		}

		if direction*newDirection < 0 {
			if secondChance {
				return false
			}
			secondChance = true
			continue
		}
		direction = newDirection
	}

	return true
}

func main() {
	startTime := time.Now()
	fmt.Print("Reading input file...")

	f, err := os.Open("02.1/input.txt")
	if err != nil {
		fmt.Println("Error Opening File")
		os.Exit(1)
	}
	fmt.Print("OK\n")

	defer f.Close()

	scanner := bufio.NewScanner(f)

	fmt.Print("Parsing File data...")
	safeReports := 0
	for scanner.Scan() {
		report := strings.Split(scanner.Text(), " ")
		if isSafe(report) {
			safeReports++
		} else if isSafe(reverse(report)) {
			safeReports++
		}
	}
	fmt.Print("OK\n")

	fmt.Println(fmt.Sprint("Number of safe reports: ", safeReports))

	endTime := time.Now()
	fmt.Println(fmt.Sprint("Program executed in ", endTime.Sub(startTime)))
}
