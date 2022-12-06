package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	input, _ := os.ReadFile("input.txt")
	sInput := strings.Split(string(input), "\n\n")

	part1 := strings.Split(sInput[0], "\n")
	part2 := strings.Split(sInput[1], "\n")

	var stacks [9][]string

	for i := 0; i < 8; i++ {
		count := 0
		for j := 1; j < 36; j = j + 4 {
			if string(part1[i][j]) != " " {
				stacks[count] = append(stacks[count], string(part1[i][j]))
			}
			count++
		}
	}

	for i := 0; i < 9; i++ {
		stacks[i] = reverse(stacks[i])
	}

	temp := stacks[0]

	for i := 0; i < len(part2); i++ {
		inst := strings.Fields(part2[i])
		index, _ := strconv.Atoi(inst[1])
		from, _ := strconv.Atoi(inst[3])
		to, _ := strconv.Atoi(inst[5])
		//Part 1
		//stacks[from-1], stacks[to-1] = move(index, stacks[from-1], stacks[to-1])
		//Part 2
		temp, stacks[from-1] = remove2(stacks[from-1], index)
		stacks[to-1] = add(stacks[to-1], temp)
	}

	var res []string

	for i := 0; i < 9; i++ {
		res = append(res, stacks[i][len(stacks[i])-1])
	}

	//fmt.Println("Part 1: ", res)
	fmt.Println("Part 2: ", res)

}

func remove(stack []string) ([]string, []string) {
	slen := len(stack)
	temp := stack[slen-1 : slen]
	stack = stack[:slen-1]
	return temp, stack
}

// Part 2
func remove2(stack []string, index int) ([]string, []string) {
	slen := len(stack)
	temp := stack[slen-index : slen]
	stack = stack[:slen-index]
	return temp, stack
}

func add(stack []string, temp []string) []string {
	stack = append(stack, temp...)
	return stack
}

func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func move(amount int, from []string, to []string) ([]string, []string) {
	temp := from

	for i := 0; i < amount; i++ {
		temp, from = remove(from)
		to = add(to, temp)
	}
	return from, to
}
