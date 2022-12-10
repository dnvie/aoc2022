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

	r, rcount, x := make([]int, 0), 0, 1
	r2, line, rcount2, x2 := make([]string, 0), genLine(), 0, 1
	for _, s := range input {
		line = checkLine(line, x2)
		ins := strings.Split(s, " ")
		if ins[0] == "noop" {
			r = append(r, x)
			rcount++
			rcount2++
			r2 = drawPixel(line, ((rcount2 - 1) % 40), r2)
		} else {
			val, _ := strconv.Atoi(ins[1])
			rcount++
			r = append(r, x)
			rcount++
			r = append(r, x)
			x += val
			rcount2++
			r2 = drawPixel(line, ((rcount2 - 1) % 40), r2)
			rcount2++
			r2 = drawPixel(line, ((rcount2 - 1) % 40), r2)
			x2 += val
			line = checkLine(line, x2)
		}
	}

	res1 := 1
	for i := 20; i <= 220; i += 40 {
		res1 += r[i-1] * i
	}
	fmt.Println("Part 1:", res1-1)

	fmt.Println("Part 2:")
	for j := 0; j < 240; j++ {
		fmt.Print(r2[j])
		if j == 40-1 || j == 80-1 || j == 120-1 || j == 160-1 || j == 200-1 || j == 240-1 {
			fmt.Println("")
		}
	}
}

func drawPixel(line []string, rcount2 int, r2 []string) []string {
	if inSprite(line, rcount2) {
		r2 = append(r2, "█")
	} else {
		r2 = append(r2, " ")
	}
	return r2
}

func checkLine(line []string, x2 int) []string {
	line = genLine()
	if x2 < 39 && 0 < x2 {
		line[x2-1], line[x2], line[x2+1] = "#", "#", "#"
	}
	return line
}

func genLine() []string {
	temp := make([]string, 0)
	for i := 0; i <= 40; i++ {
		temp = append(temp, ".")
	}
	return temp
}

func inSprite(line []string, pos int) bool {
	indices := make([]int, 0)
	for i, s := range line {
		if s == "#" {
			indices = append(indices, i)
		}
	}

	if len(indices) == 3 {
		if indices[0] == pos || indices[1] == pos || indices[2] == pos {
			return true
		}
	}
	return false
}
