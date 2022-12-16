package main

import (
	"fmt"
	"os"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	res := eval(string(input), 4)
	res2 := eval(string(input), 14)
	fmt.Println("Part 1:", res)
	fmt.Println("Part 2:", res2)
}

func eval(input string, amount int) int {
	ilen := len(input)

	for i := 0; i <= ilen-amount; i++ {
		temp := input[i : i+amount]
		map1 := make(map[string]int)
		for j := 0; j < amount; j++ {
			map1[string(temp[j])] = 0
		}
		if len(map1) == amount {
			return i + amount
		}
	}
	return 0
}
