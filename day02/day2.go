package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	res := 0
	res2 := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		res += eval(words[0], words[1])
		res2 += eval2(words[0], words[1])
	}

	fmt.Println("Part 1: ", res)
	fmt.Println("Part 2: ", res2)
}

func eval(p1 string, p2 string) int {
	if p1 == "A" {
		if p2 == "X" {
			return 4
		} else if p2 == "Y" {
			return 8
		} else {
			return 3
		}
	} else if p1 == "B" {
		if p2 == "X" {
			return 1
		} else if p2 == "Y" {
			return 5
		} else {
			return 9
		}
	} else {
		if p2 == "X" {
			return 7
		} else if p2 == "Y" {
			return 2
		} else {
			return 6
		}
	}
}

func eval2(p1 string, p2 string) int {

	if p1 == "A" {
		if p2 == "X" {
			return 3
		} else if p2 == "Y" {
			return 4
		} else {
			return 8
		}
	} else if p1 == "B" {
		if p2 == "X" {
			return 1
		} else if p2 == "Y" {
			return 5
		} else {
			return 9
		}
	} else {
		if p2 == "X" {
			return 2
		} else if p2 == "Y" {
			return 6
		} else {
			return 7
		}
	}
}
