package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	i, _ := os.ReadFile("input.txt")
	input := strings.Split(string(i), "\n")
	allpoints := make([]Point, 0)

	for i := 0; i < len(input); i++ {
		in := strings.Split(strings.ReplaceAll(input[i], "-", ""), ">")
		points := make([]Point, 0)
		for i := 0; i < len(in); i++ {
			values := strings.Fields(in[i])
			var i, j int
			fmt.Sscanf(values[i], "%d,%d", &j, &i)
			points = append(points, Point{i: i, j: j})
		}

		for i := 0; i < len(points)-1; i++ {
			allpoints = append(allpoints, getRange(points[i].i, points[i].j, points[i+1].i, points[i+1].j)...)
		}
	}

	grid := make([][]string, getMaxY(allpoints)+2)
	for i := range grid {
		grid[i] = make([]string, 9999)
	}
	fillGrid(grid)

	for _, s := range allpoints {
		grid[s.i][s.j] = "#"
	}

	p1 := true
	res1 := 0
	res2 := 0
	grain := Point{i: 0, j: 500}
	for {
		addGrid(grid, drop(grid, grain))
		if !checkBottom(grid) {
			p1 = false
		}
		if p1 {
			res1++
		}
		res2++
		if grid[0][500] == "o" {
			break
		}
	}
	fmt.Println("Part 1:", res1)
	fmt.Println("Part 2:", res2)

}

func checkBottom(grid [][]string) bool {
	for _, s := range grid[len(grid)-1] {
		if s != "." {
			if s != "#" {
				return false
			}
		}
	}
	return true
}

func drop(grid [][]string, point Point) Point {
	p := point

	if p.i >= len(grid)-1 || p.j-1 == -1 || p.j+1 > len(grid[0]) {
		return p
	}

	if grid[p.i+1][p.j] == "#" || grid[p.i+1][p.j] == "o" {
		if grid[p.i+1][p.j-1] == "#" || grid[p.i+1][p.j-1] == "o" {
			if grid[p.i+1][p.j+1] == "#" || grid[p.i+1][p.j+1] == "o" {
				return p
			} else {
				p = drop(grid, Point{i: p.i + 1, j: p.j + 1})
			}
		} else {
			p = drop(grid, Point{i: p.i + 1, j: p.j - 1})
		}
	} else {
		p = drop(grid, Point{i: p.i + 1, j: p.j})
	}
	return p
}

func addGrid(g [][]string, p Point) {
	g[p.i][p.j] = "o"
}

func getRange(x int, y int, x2 int, y2 int) []Point {
	temp := make([]Point, 0)

	if x == x2 {
		if y < y2 {
			for i := y; i <= y2; i++ {
				temp = append(temp, Point{i: x, j: i})
			}
		} else {
			for i := y2; i <= y; i++ {
				temp = append(temp, Point{i: x, j: i})
			}
		}
	} else {
		if x < x2 {
			for i := x; i <= x2; i++ {
				temp = append(temp, Point{i: i, j: y})
			}
		} else {
			for i := x2; i <= x; i++ {
				temp = append(temp, Point{i: i, j: y})
			}
		}
	}
	return temp
}

func fillGrid(grid [][]string) {
	for i := range grid {
		for j := range grid[i] {
			grid[i][j] = "."
		}
	}
}

func getMaxY(points []Point) int {
	y := 0
	for _, s := range points {
		if s.i > y {
			y = s.i
		}
	}
	return y
}

type Point struct {
	i int
	j int
}
