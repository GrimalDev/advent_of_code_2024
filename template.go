package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	startTime := time.Now()

	input, err := os.Open("08.1/input.txt")
	if err != nil {
		fmt.Println("Error Opening File")
		os.Exit(1)
	}

	defer input.Close()

	scanner := bufio.NewScanner(input)

	var result int

	for scanner.Scan() {
	}

	fmt.Println("The result is:", result)

	fmt.Println(fmt.Sprint("Script took: ", time.Now().Sub(startTime)))
}
