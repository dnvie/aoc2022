package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	filename := "input.txt"
	readInput(filename)
}

func readInput(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var values = []int{}
	curVal := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			values = append(values, curVal)
			curVal = 0
		} else {
			intVar, err := strconv.Atoi((input))
			if err != nil {
				fmt.Println("Error converting to int")
			}
			curVal += intVar
		}
	}

	sort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})

	fmt.Println("Part 1: ", values[len(values)-1])
	fmt.Println("Part 2: ", values[len(values)-1]+values[len(values)-2]+values[len(values)-3])
}
