package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	i, _ := os.ReadFile("input.txt")
	input := strings.Split(string(i), "\n")

	arr := make([][]int, len(input))
	for i := range arr {
		arr[i] = make([]int, len(input[i]))
	}

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			temp, _ := strconv.Atoi(string(input[i][j]))
			arr[i][j] = temp
		}
	}

	res, res2 := 0, 0
	for i, s := range arr {
		for j := range s {
			if eval(arr, i, j, false) != 4 {
				res++
			}
			temp := eval(arr, i, j, true)
			if temp > res2 {
				res2 = temp
			}
		}
	}
	fmt.Println("Part 1:", res, "\nPart 2:", res2)
}

func eval(arr [][]int, i int, j int, p2 bool) int {
	count := 0
	val := arr[i][j]
	c1, c2, c3, c4 := 0, 0, 0, 0

	if j > 0 {
		for a := j - 1; a >= 0; a-- { //to left
			if p2 {
				if arr[i][a] < val {
					c1++
				} else if arr[i][a] >= val {
					c1++
					break
				}
			}
			if arr[i][a] >= val {
				count++
				break
			}
		}
		for a := i - 1; a >= 0; a-- { //to top
			if p2 {
				if arr[a][j] < val {
					c3++
				} else if arr[a][j] >= val {
					c3++
					break
				}
			}
			if arr[a][j] >= val {
				count++
				break
			}
		}
	}

	if j < len(arr) { //to right
		for a := j + 1; a < len(arr); a++ {
			if p2 {
				if arr[i][a] < val {
					c2++
				} else if arr[i][a] >= val {
					c2++
					break
				}
			}
			if arr[i][a] >= val {
				count++
				break
			}
		}
	}

	if i < len(arr) { //to bottom
		for a := i + 1; a < len(arr); a++ {
			if p2 {
				if arr[a][j] < val {
					c4++
				} else if arr[a][j] >= val {
					c4++
					break
				}
			}
			if arr[a][j] >= val {
				count++
				break
			}
		}
	}
	if p2 {
		return c1 * c2 * c3 * c4
	}
	return count
}
