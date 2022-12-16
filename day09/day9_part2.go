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
			arr[i][j] = make([]string, 10)
		}
	}

	arr = fillArr(arr)
	var points []Point
	knots := []string{"H", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for i := 0; i < 10; i++ {
		points = append(points, Point{i: maxH / 2, j: maxW / 2})
		arr[maxH/2][maxW/2][i] = knots[i]
	}

	for _, s := range input {
		points[0].i, points[0].j, points[1].i, points[1].j = moveL0(arr, strings.Split(s, " "), points[0].i, points[0].j, points)
	}

	res2 := 0
	for _, s := range arr {
		for _, x := range s {
			if x[9] == "8" {
				res2++
			}
		}
	}

	fmt.Println("Part 2:", res2+1)
}

type Point struct {
	i int
	j int
}

func move(arr [][][]string, hi int, hj, ti int, tj int, l int, s string) (x int, y int) {
	if hi == ti+2 && hj == tj {
		arr[ti+1][tj][l] = s
		return ti + 1, tj
	}

	if hi == ti-2 && hj == tj {
		arr[ti-1][tj][l] = s
		return ti - 1, tj
	}

	if hi == ti && hj == tj+2 {
		arr[ti][tj+1][l] = s
		return ti, tj + 1
	}

	if hi == ti && hj == tj-2 {
		arr[ti][tj-1][l] = s
		return ti, tj - 1
	}

	if (hi == ti+2 && hj == tj+2) || (hi == ti+2 && hj == tj+1) || (hi == ti+1 && hj == tj+2) {
		arr[ti+1][tj+1][l] = s
		return ti + 1, tj + 1
	}

	if (hi == ti-1 && hj == tj+2) || (hi == ti-2 && hj == tj+2) || (hi == ti-2 && hj == tj+1) {
		arr[ti-1][tj+1][l] = s
		return ti - 1, tj + 1
	}

	if (hi == ti-1 && hj == tj-2) || (hi == ti-2 && hj == tj-1) || (hi == ti-2 && hj == tj-2) {
		arr[ti-1][tj-1][l] = s
		return ti - 1, tj - 1
	}

	if (hi == ti+2 && hj == tj-1) || (hi == ti+2 && hj == tj-2) || (hi == ti+1 && hj == tj-2) {
		arr[ti+1][tj-1][l] = s
		return ti + 1, tj - 1
	}

	return ti, tj
}

func moveAll(arr [][][]string, p []Point) {
	for pc := 1; pc < 9; pc++ {
		s := strconv.Itoa(pc)
		if !checkDist(arr, p[pc].i, p[pc].j, p[pc+1].i, p[pc+1].j) {
			p[pc+1].i, p[pc+1].j = move(arr, p[pc].i, p[pc].j, p[pc+1].i, p[pc+1].j, pc+1, s)
		}
	}
}

func moveL0(arr [][][]string, input []string, i int, j int, p []Point) (x int, y int, x1 int, y1 int) {
	dir := input[0]
	num, _ := strconv.Atoi(input[1])

	switch dir {
	case "R":
		for x := j; x <= num+j; x++ {
			arr[i][x][0] = "H"
			if !checkDist(arr, i, x, p[1].i, p[1].j) {
				p[1].i, p[1].j = move(arr, i, x, p[1].i, p[1].j, 1, "1")
			}
			moveAll(arr, p)
		}
		return i, j + num, p[1].i, p[1].j

	case "L":
		for x := j; x >= j-num; x-- {
			arr[i][x][0] = "H"
			if !checkDist(arr, i, x, p[1].i, p[1].j) {
				p[1].i, p[1].j = move(arr, i, x, p[1].i, p[1].j, 1, "1")
			}
			moveAll(arr, p)
		}
		return i, j - num, p[1].i, p[1].j

	case "U":
		for x := i; x >= i-num; x-- {
			arr[x][j][0] = "H"
			if !checkDist(arr, x, j, p[1].i, p[1].j) {
				p[1].i, p[1].j = move(arr, x, j, p[1].i, p[1].j, 1, "1")
			}
			moveAll(arr, p)
		}
		return i - num, j, p[1].i, p[1].j

	case "D":
		for x := i; x <= num+i; x++ {
			arr[x][j][0] = "H"
			if !checkDist(arr, x, j, p[1].i, p[1].j) {
				p[1].i, p[1].j = move(arr, x, j, p[1].i, p[1].j, 1, "1")
			}
			moveAll(arr, p)
		}
		return i + num, j, p[1].i, p[1].j
	}

	return i, j, p[1].i, p[1].j
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
			for k := range arr[i][j] {
				arr[i][j][k] = "."
			}
		}
	}
	return arr
}
