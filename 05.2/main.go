package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func moveItem(slice []string, from int, to int) []string {
	if to < 0 {
		to = 0
	}

	item := slice[from]
	slice = append(slice[:from], slice[from+1:]...)

	slice = append(slice[:to], append([]string{item}, slice[to:]...)...)

	return slice
}

func main() {
	startTime := time.Now()

	input, err := os.Open("05.2/input.txt")
	if err != nil {
		fmt.Println("Error Opening File")
		os.Exit(1)
	}

	defer input.Close()

	scanner := bufio.NewScanner(input)

	rules := make(map[string]int)
	var middleSum int
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "|") {
			rules[line] = 0
			continue
		}

		if line == "" {
			continue
		}

		update := strings.Split(line, ",")
		safeUpdate := true
		for i := 0; i < len(update); i++ {
			for j := i + 1; j < len(update); j++ {
				if _, ok := rules[update[j]+"|"+update[i]]; ok {
					moveItem(update, j, i)
					i = 0
					safeUpdate = false
				}
			}
		}

		if !safeUpdate {
			num, _ := strconv.Atoi(update[len(update)/2])
			middleSum += num
		}
	}

	fmt.Println("Sum of middle numbers from safe updates is: ", middleSum)

	fmt.Println(fmt.Sprint("Script took: ", time.Now().Sub(startTime)))
}
