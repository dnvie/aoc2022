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

	maxH, maxW := 1000, 1000
	arr := make([][][]string, maxH)
	for i := range arr {
		arr[i] = make([][]string, maxW)
		for j := range arr[i] {
			arr[i][j] = make([]string, 2)
		}
	}

	pi1, pj1, pi2, pj2 := maxH/2, maxW/2, maxH/2, maxW/2
	arr = fillArr(arr)
	arr[maxH/2][maxW/2][0] = "H"
	arr[maxH/2][maxW/2][1] = "T"

	for _, s := range input {
		pi1, pj1, pi2, pj2 = moveL0(arr, strings.Split(s, " "), pi1, pj1, pi2, pj2)
	}

	res1 := 0
	for _, s := range arr {
		for _, x := range s {
			if x[1] == "T" {
				res1++
			}
		}
	}
	fmt.Println("Part 1:", res1)
}

func moveL1(arr [][][]string, hi int, hj int, ti int, tj int) (x int, y int) {
	if hi == ti+2 && hj == tj {
		arr[ti+1][tj][1] = "T"
		return ti + 1, tj
	}

	if hi == ti-2 && hj == tj {
		arr[ti-1][tj][1] = "T"
		return ti - 1, tj
	}

	if hi == ti && hj == tj+2 {
		arr[ti][tj+1][1] = "T"
		return ti, tj + 1
	}

	if hi == ti && hj == tj-2 {
		arr[ti][tj-1][1] = "T"
		return ti, tj - 1
	}

	if (hi == ti+2 && hj == tj+2) || (hi == ti+2 && hj == tj+1) || (hi == ti+1 && hj == tj+2) {
		arr[ti+1][tj+1][1] = "T"
		return ti + 1, tj + 1
	}

	if (hi == ti-1 && hj == tj+2) || (hi == ti-2 && hj == tj+2) || (hi == ti-2 && hj == tj+1) {
		arr[ti-1][tj+1][1] = "T"
		return ti - 1, tj + 1
	}

	if (hi == ti-1 && hj == tj-2) || (hi == ti-2 && hj == tj-1) || (hi == ti-2 && hj == tj-2) {
		arr[ti-1][tj-1][1] = "T"
		return ti - 1, tj - 1
	}

	if (hi == ti+2 && hj == tj-1) || (hi == ti+2 && hj == tj-2) || (hi == ti+1 && hj == tj-2) {
		arr[ti+1][tj-1][1] = "T"
		return ti + 1, tj - 1
	}

	return ti, tj
}

func moveL0(arr [][][]string, input []string, i int, j int, ti int, tj int) (x int, y int, x1 int, y1 int) {
	dir := input[0]
	num, _ := strconv.Atoi(input[1])

	switch dir {
	case "R":
		for x := j; x <= num+j; x++ {
			arr[i][x][0] = "H"
			if !checkDist(arr, i, x, ti, tj) {
				ti, tj = moveL1(arr, i, x, ti, tj)
			}
		}
		return i, j + num, ti, tj

	case "L":
		for x := j; x >= j-num; x-- {
			arr[i][x][0] = "H"
			if !checkDist(arr, i, x, ti, tj) {
				ti, tj = moveL1(arr, i, x, ti, tj)
			}
		}
		return i, j - num, ti, tj

	case "U":
		for x := i; x >= i-num; x-- {
			arr[x][j][0] = "H"
			if !checkDist(arr, x, j, ti, tj) {
				ti, tj = moveL1(arr, x, j, ti, tj)
			}
		}
		return i - num, j, ti, tj

	case "D":
		for x := i; x <= num+i; x++ {
			arr[x][j][0] = "H"
			if !checkDist(arr, x, j, ti, tj) {
				ti, tj = moveL1(arr, x, j, ti, tj)
			}
		}
		return i + num, j, ti, tj
	}
	return i, j, ti, tj
}

func checkDist(arr [][][]string, i1 int, j1 int, i2 int, j2 int) bool {
	if (i2 == i1 || i2 == i1-1 || i2 == i1+1) && (j2 == j1 || j2 == j1-1 || j2 == j1+1) {
		return true
	}
	return false
}

func fillArr(arr [][][]string) [][][]string {
	for i := range arr {
		for j := range arr[i] {
			arr[i][j][0] = "."
			arr[i][j][1] = "."
		}
	}
	return arr
}
