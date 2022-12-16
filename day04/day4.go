package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	res := 0
	res2 := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		values := strings.FieldsFunc(scanner.Text(), Split)
		res += eval1(values)
		res2 += eval2(values)
	}
	fmt.Println("Result Part 1: ", res)
	fmt.Println("Result Part 2: ", res2)
}

func Split(r rune) bool {
	return r == '-' || r == ','
}

func eval1(values []string) int {
	v1, _ := strconv.Atoi(values[0])
	v2, _ := strconv.Atoi(values[1])
	v3, _ := strconv.Atoi(values[2])
	v4, _ := strconv.Atoi(values[3])
	if (v1 <= v3 && v2 >= v4) || (v3 <= v1 && v4 >= v2) {
		return 1
	}
	return 0
}

func eval2(values []string) int {
	v1, _ := strconv.Atoi(values[0])
	v2, _ := strconv.Atoi(values[1])
	v3, _ := strconv.Atoi(values[2])
	v4, _ := strconv.Atoi(values[3])
	map1 := make(map[int]int)
	for i := v1; i <= v2; i++ {
		map1[i] = i
	}

	for i := v3; i <= v4; i++ {
		if _, contains := map1[i]; contains {
			return 1
		}
	}
	return 0
}
