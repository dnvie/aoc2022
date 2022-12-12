package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	i, _ := os.ReadFile("input.txt")
	input := strings.Split(string(i), "\n\n")

	monkeymap, monkeymap2 := make(map[int]*monkey), make(map[int]*monkey)
	for i, s := range input {
		lines := strings.Split(s, "\n")
		temp, temp2 := createMonkey(lines), createMonkey(lines)
		monkeymap[i], monkeymap2[i] = temp, temp2
	}

	solve(monkeymap, false)
	solve(monkeymap2, true)
}

func solve(monkeymap map[int]*monkey, p2 bool) {
	rounds := 20
	if p2 {
		rounds = 10000
	}
	for i := 0; i < rounds; i++ {
		for i := 0; i < len(monkeymap); i++ {
			m := monkeymap[i]
			for i := range m.items {
				m.inspected++
				newLevel := eval(m.operation, m.items[i])
				if p2 {
					mod := 1
					for x := range monkeymap {
						mod *= monkeymap[x].test
					}
					newLevel = newLevel % mod
				} else {
					newLevel = newLevel / 3
				}
				nextMonkey := &monkey{}
				if newLevel%m.test == 0 {
					nextMonkey = monkeymap[m.iftrue]
				} else {
					nextMonkey = monkeymap[m.iffalse]
				}
				addItems(nextMonkey, newLevel)
				monkeymap[nextMonkey.id] = nextMonkey
				m.items[i] = -1
			}
			rearrangeItems(m)
			monkeymap[i] = m
		}
	}

	val1 := 0
	for _, s := range monkeymap {
		if s.inspected > val1 {
			val1 = s.inspected
		}
	}

	val2 := 0
	for _, s := range monkeymap {
		if s.inspected > val2 && s.inspected != val1 {
			val2 = s.inspected
		}
	}
	if p2 {
		fmt.Println("Part 2:", val1*val2)
	} else {
		fmt.Println("Part 1:", val1*val2)
	}
}

func addItems(m *monkey, i int) {
	m.items = append(m.items, i)
}

func rearrangeItems(m *monkey) {
	temp := make([]int, 0)
	for _, s := range m.items {
		if s != -1 {
			temp = append(temp, s)
		}
	}
	m.items = temp
}

func eval(o string, l int) int {
	op := strings.Split(o, " ")[0]
	val := l
	if !(strings.Split(o, " ")[1] == "old") {
		val, _ = strconv.Atoi(strings.Split(o, " ")[1])
	}
	if op == "*" {
		return l * val
	}
	return l + val
}

func createMonkey(lines []string) *monkey {
	temp := &monkey{}
	name := strings.Fields(lines[0])
	lines[1] = strings.Replace(lines[1], ",", "", -1)
	sitems := strings.Fields(lines[1])[2:]
	items := make([]int, 0)
	for _, s := range sitems {
		temp, _ := strconv.Atoi(s)
		items = append(items, temp)
	}
	operation := strings.Fields(lines[2])[4] + " " + strings.Fields(lines[2])[5]
	test, _ := strconv.Atoi(strings.Fields(lines[3])[3])
	iftrue, _ := strconv.Atoi(strings.Fields(lines[4])[5])
	iffalse, _ := strconv.Atoi(strings.Fields(lines[5])[5])
	temp.id, _ = strconv.Atoi(name[1][:len(name[1])-1])
	temp.items, temp.operation, temp.test, temp.iftrue, temp.iffalse, temp.inspected = items, operation, test, iftrue, iffalse, 0
	return temp
}

type monkey struct {
	id        int
	items     []int
	operation string
	test      int
	iftrue    int
	iffalse   int
	inspected int
}
