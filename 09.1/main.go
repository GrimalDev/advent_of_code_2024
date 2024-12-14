package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	startTime := time.Now()

	input, err := os.ReadFile("09.1/input.txt")
	if err != nil {
		fmt.Println("Error Opening File")
		os.Exit(1)
	}

	var disk []int

	var id int
	for i, n := range input {
		if n == '\n' {
			break
		}
		id2 := id
		if i%2 == 1 {
			id2 = -1
			id++
		}
		nint, _ := strconv.Atoi(string(n))
		for range nint {
			disk = append(disk, id2)
		}
	}

	for i := 0; i < len(disk); i++ {
		if disk[i] == -1 {
			for j := len(disk) - 1; j > i; j-- {
				if disk[j] != -1 {
					disk[i] = disk[j]
					disk[j] = -1
					i = 0
					break
				}
			}
		}
	}

	var flag int

	for p, id := range disk {
		if id != -1 {
			flag += p * id
		}
	}

	fmt.Println("Result is:", flag)

	fmt.Println(fmt.Sprint("Script took: ", time.Now().Sub(startTime)))
}
