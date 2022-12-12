package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	i, _ := os.ReadFile("input.txt")
	input := strings.Split(string(i), "\n")

	start, end := &Point{i: 0, j: 0, value: 1, minDist: 0, visited: false}, &Point{i: 0, j: 0, value: 26, minDist: 9999, visited: false}

	grid := make([][]*Point, len(input))
	for i := range grid {
		grid[i] = make([]*Point, len(input[0]))
	}

	p2points := make([]*Point, 0)

	for i := range grid {
		temp := strings.Split(input[i], "")
		for j := range grid[i] {
			if getNum(temp[j]) == -2 {
				start.i, start.j = i, j
				grid[i][j] = start
				p2points = append(p2points, start)
			} else if getNum(temp[j]) == -3 {
				end.i, end.j = i, j
				grid[i][j] = end
			} else {
				if getNum(temp[j]) == 1 {
					grid[i][j] = &Point{i: i, j: j, value: getNum(temp[j]), minDist: 99, visited: false}
					p2points = append(p2points, grid[i][j])
				} else {
					grid[i][j] = &Point{i: i, j: j, value: getNum(temp[j]), minDist: 9999, visited: false}
				}
			}
		}
	}

	points := make([]*Point, 0)
	points = append(points, start)

	for i := 0; i < len(points); i++ {
		points = getDis(grid, points[i].i, points[i].j, points)
	}
	fmt.Println("Part 1:", end.minDist)

	for _, s := range p2points {
		s.minDist = 0
	}

	res2 := end.minDist
	for i := 0; i < len(p2points); i++ {
		points2 := make([]*Point, 0)
		points2 = append(points2, p2points[i])

		for i := 0; i < len(points2); i++ {
			points2 = getDis(grid, points2[i].i, points2[i].j, points2)
		}
		if end.minDist < res2 {
			res2 = end.minDist
		}
	}
	fmt.Println("Part 2:", end.minDist)

}

func getDis(grid [][]*Point, i int, j int, points []*Point) []*Point {
	if inBounds(grid, i-1, j) {
		if !(grid[i-1][j].visited) {
			if grid[i-1][j].minDist > grid[i][j].minDist+1 {
				if grid[i-1][j].value <= grid[i][j].value+1 {
					grid[i-1][j].minDist = grid[i][j].minDist + 1
					points = append(points, grid[i-1][j])
				}
			}
		}
	}

	if inBounds(grid, i+1, j) {
		if !(grid[i+1][j].visited) {
			if grid[i+1][j].minDist > grid[i][j].minDist+1 {
				if grid[i+1][j].value <= grid[i][j].value+1 {
					grid[i+1][j].minDist = grid[i][j].minDist + 1
					points = append(points, grid[i+1][j])
				}
			}
		}
	}

	if inBounds(grid, i, j-1) {
		if !(grid[i][j-1].visited) {
			if grid[i][j-1].minDist > grid[i][j].minDist+1 {
				if grid[i][j-1].value <= grid[i][j].value+1 {
					grid[i][j-1].minDist = grid[i][j].minDist + 1
					points = append(points, grid[i][j-1])
				}
			}
		}
	}

	if inBounds(grid, i, j+1) {
		if !(grid[i][j+1].visited) {
			if grid[i][j+1].minDist > grid[i][j].minDist+1 {
				if grid[i][j+1].value <= grid[i][j].value+1 {
					grid[i][j+1].minDist = grid[i][j].minDist + 1
					points = append(points, grid[i][j+1])
				}
			}
		}
	}
	return points
}

func inBounds(grid [][]*Point, i int, j int) bool {
	if i < 0 || j < 0 || len(grid) <= i || len(grid[0]) <= j {
		return false
	}
	return true
}

func getNum(s string) int {
	lowercase := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	for i, c := range lowercase {
		if s == "S" {
			return -2
		}
		if s == "E" {
			return -3
		}
		if c == s {
			return i + 1
		}
	}
	return -1
}

type Point struct {
	i       int
	j       int
	value   int
	minDist int
	visited bool
}
