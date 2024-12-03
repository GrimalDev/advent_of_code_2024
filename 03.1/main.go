package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	startTime := time.Now()
	fmt.Print("Reading input file...")

	f, err := os.Open("03.1/input.txt")
	if err != nil {
		fmt.Println("Error Opening File")
		os.Exit(1)
	}
	fmt.Print("OK\n")

	defer f.Close()

	scanner := bufio.NewScanner(f)

	fmt.Print("Searching the Pattern...")
	r, err := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)`)
	if err != nil {
		fmt.Print("Error searching the Pattern")
		os.Exit(1)
	}

	var matches [][]string
	for scanner.Scan() {
		matches = append(matches, r.FindAllStringSubmatch(scanner.Text(), -1)...)
	}

	var multiplication int
	for _, match := range matches {
		n1, _ := strconv.Atoi(match[1])
		n2, _ := strconv.Atoi(match[2])

		multiplication += n1 * n2
	}
	fmt.Print("OK\n")

	fmt.Println(fmt.Sprint("The multiplication result is: ", multiplication))
	fmt.Println(fmt.Sprint("Script took: ", time.Now().Sub(startTime)))
}
