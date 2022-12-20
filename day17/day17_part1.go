package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	i, _ := os.ReadFile("input.txt")

	blocked := make([][]bool, 4000)
	for i := range blocked {
		blocked[i] = make([]bool, 7)
	}
	fillGrid(blocked)

	direction := 0
	for rocks := 1; rocks <= 2022; rocks++ {
		r := createRock(getHighest(blocked), rocks)
		for !r.resting {
			r = moveRock(string(i[direction%10091]), r, blocked, rocks)
			direction++
		}
	}

	fmt.Println(4000 - getHighest(blocked) - 4)
	fmt.Println("Time elapsed:", time.Since(start))

}

func fillGrid(grid [][]bool) {
	for i := range grid {
		for j := range grid[i] {
			grid[i][j] = false
		}
	}
}

func createRock(h int, t int) Rock {
	if t%5 == 1 {
		p1 := Point{i: h, j: 2}
		p2 := Point{i: h, j: 3}
		p3 := Point{i: h, j: 4}
		p4 := Point{i: h, j: 5}
		return Rock{p1: p1, p2: p2, p3: p3, p4: p4}
	} else if t%5 == 2 {
		p1 := Point{i: h - 2, j: 3}
		p2 := Point{i: h - 1, j: 2}
		p3 := Point{i: h - 1, j: 3}
		p4 := Point{i: h - 1, j: 4}
		p5 := Point{i: h, j: 3}
		return Rock{p1: p1, p2: p2, p3: p3, p4: p4, p5: p5}
	} else if t%5 == 3 {
		p1 := Point{i: h - 2, j: 4}
		p2 := Point{i: h - 1, j: 4}
		p3 := Point{i: h, j: 2}
		p4 := Point{i: h, j: 3}
		p5 := Point{i: h, j: 4}
		return Rock{p1: p1, p2: p2, p3: p3, p4: p4, p5: p5}
	} else if t%5 == 4 {
		p1 := Point{i: h - 3, j: 2}
		p2 := Point{i: h - 2, j: 2}
		p3 := Point{i: h - 1, j: 2}
		p4 := Point{i: h, j: 2}
		return Rock{p1: p1, p2: p2, p3: p3, p4: p4}
	} else {
		p1 := Point{i: h - 1, j: 2}
		p2 := Point{i: h - 1, j: 3}
		p3 := Point{i: h, j: 2}
		p4 := Point{i: h, j: 3}
		return Rock{p1: p1, p2: p2, p3: p3, p4: p4}
	}
}

func moveRock(s string, r Rock, blocked [][]bool, t int) Rock {
	if t%5 == 1 {
		if s == ">" {
			if r.p4.j < 6 && blocked[r.p4.i][r.p4.j+1] == false {
				r.p1.j++
				r.p2.j++
				r.p3.j++
				r.p4.j++
			}
		} else {
			if r.p1.j > 0 && blocked[r.p1.i][r.p1.j-1] == false {
				r.p1.j--
				r.p2.j--
				r.p3.j--
				r.p4.j--
			}
		}
		if r.p1.i+1 < len(blocked) {
			if blocked[r.p1.i+1][r.p1.j] == false && blocked[r.p2.i+1][r.p2.j] == false && blocked[r.p3.i+1][r.p3.j] == false && blocked[r.p4.i+1][r.p4.j] == false {
				r.p1.i++
				r.p2.i++
				r.p3.i++
				r.p4.i++
			} else {
				blocked[r.p1.i][r.p1.j] = true
				blocked[r.p2.i][r.p2.j] = true
				blocked[r.p3.i][r.p3.j] = true
				blocked[r.p4.i][r.p4.j] = true
				r.resting = true
			}
		} else {
			blocked[r.p1.i][r.p1.j] = true
			blocked[r.p2.i][r.p2.j] = true
			blocked[r.p3.i][r.p3.j] = true
			blocked[r.p4.i][r.p4.j] = true
			r.resting = true
		}
		return r
	} else if t%5 == 2 {
		if s == ">" {
			if r.p4.j < 6 && blocked[r.p1.i][r.p1.j+1] == false && blocked[r.p4.i][r.p4.j+1] == false && blocked[r.p5.i][r.p5.j+1] == false {
				r.p1.j++
				r.p2.j++
				r.p3.j++
				r.p4.j++
				r.p5.j++
			}
		} else {
			if r.p2.j > 0 && blocked[r.p1.i][r.p1.j-1] == false && blocked[r.p2.i][r.p2.j-1] == false && blocked[r.p5.i][r.p5.j-1] == false {
				r.p1.j--
				r.p2.j--
				r.p3.j--
				r.p4.j--
				r.p5.j--
			}
		}
		if r.p5.i+1 < len(blocked) {
			if blocked[r.p2.i+1][r.p2.j] == false && blocked[r.p4.i+1][r.p4.j] == false && blocked[r.p5.i+1][r.p5.j] == false {
				r.p1.i++
				r.p2.i++
				r.p3.i++
				r.p4.i++
				r.p5.i++
			} else {
				blocked[r.p1.i][r.p1.j] = true
				blocked[r.p2.i][r.p2.j] = true
				blocked[r.p3.i][r.p3.j] = true
				blocked[r.p4.i][r.p4.j] = true
				blocked[r.p5.i][r.p5.j] = true
				r.resting = true
			}
		} else {
			blocked[r.p1.i][r.p1.j] = true
			blocked[r.p2.i][r.p2.j] = true
			blocked[r.p3.i][r.p3.j] = true
			blocked[r.p4.i][r.p4.j] = true
			blocked[r.p5.i][r.p5.j] = true
			r.resting = true
		}
		return r
	} else if t%5 == 3 {
		if s == ">" {
			if r.p1.j < 6 && blocked[r.p1.i][r.p1.j+1] == false && blocked[r.p2.i][r.p2.j+1] == false && blocked[r.p5.i][r.p5.j+1] == false {
				r.p1.j++
				r.p2.j++
				r.p3.j++
				r.p4.j++
				r.p5.j++
			}
		} else {
			if r.p3.j > 0 && blocked[r.p1.i][r.p1.j-1] == false && blocked[r.p2.i][r.p2.j-1] == false && blocked[r.p3.i][r.p3.j-1] == false {
				r.p1.j--
				r.p2.j--
				r.p3.j--
				r.p4.j--
				r.p5.j--
			}
		}
		if r.p3.i+1 < len(blocked) {
			if blocked[r.p3.i+1][r.p3.j] == false && blocked[r.p4.i+1][r.p4.j] == false && blocked[r.p5.i+1][r.p5.j] == false {
				r.p1.i++
				r.p2.i++
				r.p3.i++
				r.p4.i++
				r.p5.i++
			} else {
				blocked[r.p1.i][r.p1.j] = true
				blocked[r.p2.i][r.p2.j] = true
				blocked[r.p3.i][r.p3.j] = true
				blocked[r.p4.i][r.p4.j] = true
				blocked[r.p5.i][r.p5.j] = true
				r.resting = true
			}
		} else {
			blocked[r.p1.i][r.p1.j] = true
			blocked[r.p2.i][r.p2.j] = true
			blocked[r.p3.i][r.p3.j] = true
			blocked[r.p4.i][r.p4.j] = true
			blocked[r.p5.i][r.p5.j] = true
			r.resting = true
		}
		return r
	} else if t%5 == 4 {
		if s == ">" {
			if r.p4.j < 6 && blocked[r.p1.i][r.p1.j+1] == false && blocked[r.p2.i][r.p2.j+1] == false && blocked[r.p3.i][r.p3.j+1] == false && blocked[r.p4.i][r.p4.j+1] == false {
				r.p1.j++
				r.p2.j++
				r.p3.j++
				r.p4.j++
			}
		} else {
			if r.p1.j > 0 && blocked[r.p1.i][r.p1.j-1] == false && blocked[r.p2.i][r.p2.j-1] == false && blocked[r.p3.i][r.p3.j-1] == false && blocked[r.p4.i][r.p4.j-1] == false {
				r.p1.j--
				r.p2.j--
				r.p3.j--
				r.p4.j--
			}
		}
		if r.p4.i+1 < len(blocked) {
			if blocked[r.p4.i+1][r.p4.j] == false {
				r.p1.i++
				r.p2.i++
				r.p3.i++
				r.p4.i++
			} else {
				blocked[r.p1.i][r.p1.j] = true
				blocked[r.p2.i][r.p2.j] = true
				blocked[r.p3.i][r.p3.j] = true
				blocked[r.p4.i][r.p4.j] = true
				r.resting = true
			}
		} else {
			blocked[r.p1.i][r.p1.j] = true
			blocked[r.p2.i][r.p2.j] = true
			blocked[r.p3.i][r.p3.j] = true
			blocked[r.p4.i][r.p4.j] = true
			r.resting = true
		}
		return r
	} else {
		if s == ">" {
			if r.p4.j < 6 && blocked[r.p2.i][r.p2.j+1] == false && blocked[r.p4.i][r.p4.j+1] == false {
				r.p1.j++
				r.p2.j++
				r.p3.j++
				r.p4.j++
			}
		} else {
			if r.p1.j > 0 && blocked[r.p1.i][r.p1.j-1] == false && blocked[r.p3.i][r.p3.j-1] == false {
				r.p1.j--
				r.p2.j--
				r.p3.j--
				r.p4.j--
			}
		}
		if r.p1.i+1 < len(blocked) {
			if blocked[r.p3.i+1][r.p3.j] == false && blocked[r.p4.i+1][r.p4.j] == false {
				r.p1.i++
				r.p2.i++
				r.p3.i++
				r.p4.i++
			} else {
				blocked[r.p1.i][r.p1.j] = true
				blocked[r.p2.i][r.p2.j] = true
				blocked[r.p3.i][r.p3.j] = true
				blocked[r.p4.i][r.p4.j] = true
				r.resting = true
			}
		} else {
			blocked[r.p1.i][r.p1.j] = true
			blocked[r.p2.i][r.p2.j] = true
			blocked[r.p3.i][r.p3.j] = true
			blocked[r.p4.i][r.p4.j] = true
			r.resting = true
		}
		return r
	}
}

func getHighest(grid [][]bool) int {
	x := len(grid)

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] {
				if i < x {
					x = i
				}
			}
		}
	}

	return x - 4
}

type Rock struct {
	resting            bool
	p1, p2, p3, p4, p5 Point
}

type Point struct {
	i int
	j int
}
