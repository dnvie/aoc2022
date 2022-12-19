// Part 1 only
package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	i, _ := os.ReadFile("input.txt")
	input := strings.Split(string(i), "\n")

	grid, gridB := make([]bool, 10000000), make([]bool, 10000000)

	for _, s := range input {
		var sx, sy, bx, by int
		fmt.Sscanf(s, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		if by == 2000000 {
			gridB[bx+5000000] = true
		}
		temp := manDist(sx, sy, bx, by)
		fill(sx, sy, temp*2+1, grid)
	}

	elapsed := time.Since(start)
	fmt.Println("Part 1:", count(grid, gridB))
	fmt.Println("Time elapsed:", elapsed)
}

func manDist(x1 int, y1 int, x2 int, y2 int) int {
	return int(math.Abs(float64(x1)-float64(x2)) + math.Abs(float64(y1)-float64(y2)))
}

func checkIfOnLine(y int, d int) bool {
	if y-d <= 2000000 && y+d >= 2000000 {
		return true
	}
	return false
}

func fill(x int, y int, c int, grid []bool) []bool {
	offset := 2000000 - y
	from := (x - (c / 2)) + int(math.Abs(float64(offset)))
	to := (x + (c / 2)) - int(math.Abs(float64(offset)))

	if checkIfOnLine(y, (c / 2)) {
		for i := from; i <= to; i++ {
			grid[i+5000000] = true
		}
	}
	return grid
}

func count(grid []bool, gridB []bool) int {
	res, res2 := 0, 0

	for _, s := range grid {
		if s == true {
			res++
		}
	}

	for _, s := range gridB {
		if s == true {
			res2++
		}
	}
	return res - res2
}
