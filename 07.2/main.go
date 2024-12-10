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

func applyOperant(operant string, lastOperant string, total int, lastTotal int, num int, numAfter int) int {
	switch operant {
	case "+":
		total += numAfter
	case "*":
		total *= numAfter
	case "||":
		numsConcat, _ := strconv.Atoi(strconv.Itoa(num) + strconv.Itoa(numAfter))
		if total == lastTotal {
			total = numsConcat
		} else {
			total = applyOperant(lastOperant, lastOperant, lastTotal, lastTotal, lastTotal, numsConcat)
		}
	}
	fmt.Println(lastTotal, lastOperant, num, operant, numAfter, "=", total)
	return total
}

func main() {
	startTime := time.Now()

	input, err := os.Open("07.2/input.txt")
	if err != nil {
		fmt.Println("Error Opening File")
		os.Exit(1)
	}

	defer input.Close()

	scanner := bufio.NewScanner(input)

	var workingSum int
	operants := []string{"+", "*", "||"}

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

			var lastOperant string
			results := []int{nums[0]}
			for j := 0; j < len(currentOperants); j++ {
				result := results[len(results)-1]
				lastResult := result
				if len(results) > 1 {
					lastResult = results[len(results)-2]
				}
				results = append(results, applyOperant(currentOperants[j], lastOperant, result, lastResult, nums[j], nums[j+1]))

				lastOperant = currentOperants[j]
			}

			if results[len(results)-1] == target {
				workingSum += target
				break
			}
		}
	}

	fmt.Println("The sum is: ", workingSum)

	fmt.Println(fmt.Sprint("Script took: ", time.Now().Sub(startTime)))
}
