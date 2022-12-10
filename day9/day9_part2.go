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

	maxH, maxW := 10000, 10000

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
		points[0].i, points[0].j, points[1].i, points[1].j = moveL0(arr, strings.Split(s, " "), points[0].i, points[0].j, points[1].i, points[1].j, points)
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
	//up
	if hi == ti+2 && hj == tj {
		arr[ti+1][tj][l] = s
		return ti + 1, tj
	}

	//down
	if hi == ti-2 && hj == tj {
		arr[ti-1][tj][l] = s
		return ti - 1, tj
	}

	//left
	if hi == ti && hj == tj+2 {
		arr[ti][tj+1][l] = s
		return ti, tj + 1
	}

	//right
	if hi == ti && hj == tj-2 {
		arr[ti][tj-1][l] = s
		return ti, tj - 1
	}

	//left up to right down
	if (hi == ti+2 && hj == tj+2) || (hi == ti+2 && hj == tj+1) || (hi == ti+1 && hj == tj+2) {
		arr[ti+1][tj+1][l] = s
		return ti + 1, tj + 1
	}

	//left down to right up
	if (hi == ti-1 && hj == tj+2) || (hi == ti-2 && hj == tj+2) || (hi == ti-2 && hj == tj+1) {
		arr[ti-1][tj+1][l] = s
		return ti - 1, tj + 1
	}

	//right down left up
	if (hi == ti-1 && hj == tj-2) || (hi == ti-2 && hj == tj-1) || (hi == ti-2 && hj == tj-2) {
		arr[ti-1][tj-1][l] = s
		return ti - 1, tj - 1
	}

	//right up left down
	if (hi == ti+2 && hj == tj-1) || (hi == ti+2 && hj == tj-2) || (hi == ti+1 && hj == tj-2) {
		arr[ti+1][tj-1][l] = s
		return ti + 1, tj - 1
	}

	return ti, tj
}

func moveL0(arr [][][]string, input []string, i int, j int, ti int, tj int, p []Point) (x int, y int, x1 int, y1 int) {
	dir := input[0]
	num, _ := strconv.Atoi(input[1])

	switch dir {
	case "R":
		for x := j; x <= num+j; x++ {
			arr[i][x][0] = "H"
			if !checkDist(arr, i, x, ti, tj) {
				ti, tj = move(arr, i, x, ti, tj, 1, "1")
			}
			if !checkDist(arr, ti, tj, p[2].i, p[2].j) {
				p[2].i, p[2].j = move(arr, ti, tj, p[2].i, p[2].j, 2, "2")
			}
			for pc := 2; pc < 9; pc++ {
				s := strconv.Itoa(pc)
				if !checkDist(arr, p[pc].i, p[pc].j, p[pc+1].i, p[pc+1].j) {
					p[pc+1].i, p[pc+1].j = move(arr, p[pc].i, p[pc].j, p[pc+1].i, p[pc+1].j, pc+1, s)
				}
			}
		}

		for x := j; x < num+j; x++ {
			arr[i][x][0] = "."
		}

		return i, j + num, ti, tj

	case "L":
		for x := j; x >= j-num; x-- {
			arr[i][x][0] = "H"
			if !checkDist(arr, i, x, ti, tj) {
				ti, tj = move(arr, i, x, ti, tj, 1, "1")
			}
			if !checkDist(arr, ti, tj, p[2].i, p[2].j) {
				p[2].i, p[2].j = move(arr, ti, tj, p[2].i, p[2].j, 2, "2")
			}
			for pc := 2; pc < 9; pc++ {
				s := strconv.Itoa(pc)
				if !checkDist(arr, p[pc].i, p[pc].j, p[pc+1].i, p[pc+1].j) {
					p[pc+1].i, p[pc+1].j = move(arr, p[pc].i, p[pc].j, p[pc+1].i, p[pc+1].j, pc+1, s)
				}
			}
		}

		for x := j; x > j-num; x-- {
			arr[i][x][0] = "."
		}
		return i, j - num, ti, tj

	case "U":
		for x := i; x >= i-num; x-- {
			arr[x][j][0] = "H"
			if !checkDist(arr, x, j, ti, tj) {
				ti, tj = move(arr, x, j, ti, tj, 1, "1")
			}
			if !checkDist(arr, ti, tj, p[2].i, p[2].j) {
				p[2].i, p[2].j = move(arr, ti, tj, p[2].i, p[2].j, 2, "2")
			}
			for pc := 2; pc < 9; pc++ {
				s := strconv.Itoa(pc)
				if !checkDist(arr, p[pc].i, p[pc].j, p[pc+1].i, p[pc+1].j) {
					p[pc+1].i, p[pc+1].j = move(arr, p[pc].i, p[pc].j, p[pc+1].i, p[pc+1].j, pc+1, s)
				}
			}
		}

		for x := i; x > i-num; x-- {
			arr[x][j][0] = "."
		}
		return i - num, j, ti, tj

	case "D":
		for x := i; x <= num+i; x++ {
			arr[x][j][0] = "H"
			if !checkDist(arr, x, j, ti, tj) {
				ti, tj = move(arr, x, j, ti, tj, 1, "1")
			}
			if !checkDist(arr, ti, tj, p[2].i, p[2].j) {
				p[2].i, p[2].j = move(arr, ti, tj, p[2].i, p[2].j, 2, "2")
			}
			for pc := 2; pc < 9; pc++ {
				s := strconv.Itoa(pc)
				if !checkDist(arr, p[pc].i, p[pc].j, p[pc+1].i, p[pc+1].j) {
					p[pc+1].i, p[pc+1].j = move(arr, p[pc].i, p[pc].j, p[pc+1].i, p[pc+1].j, pc+1, s)
				}
			}
		}

		for x := i; x < num+i; x++ {
			arr[x][j][0] = "."
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

func printArr(arr [][][]string, l int) {
	for i := range arr {
		for j := range arr[i] {
			fmt.Print(arr[i][j][l], " ")
		}
		fmt.Println("")
	}
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
