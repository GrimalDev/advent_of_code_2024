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

func main() {
	startTime := time.Now()

	input, err := os.Open("07.1/input.txt")
	if err != nil {
		fmt.Println("Error Opening File")
		os.Exit(1)
	}

	defer input.Close()

	scanner := bufio.NewScanner(input)

	var workingSum int
	operants := []string{"+", "*"}

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ": ")
		target, _ := strconv.Atoi(parts[0])
		numStrs := strings.Split(parts[1], " ")
		nums := make([]int, len(numStrs))
		for i, ns := range numStrs {
			nums[i], _ = strconv.Atoi(ns)
		}

		totalCombinations := int(math.Pow(float64(len(operants)), float64(len(numStrs))))

		for i := 0; i < totalCombinations; i++ {
			combinationIndex := i
			currentOperants := make([]string, len(nums)-1)
			for j := 0; j < len(nums)-1; j++ {
				currentOperants[j] = operants[combinationIndex%len(operants)]
				combinationIndex /= len(operants)
			}

			result := nums[0]
			for j := 0; j < len(currentOperants); j++ {
				switch currentOperants[j] {
				case "+":
					result += nums[j+1]
				case "*":
					result *= nums[j+1]
				}
			}

			if result == target {
				workingSum += target
				break
			}
		}
	}

	fmt.Println("The sum is: ", workingSum)

	fmt.Println(fmt.Sprint("Script took: ", time.Now().Sub(startTime)))
}
