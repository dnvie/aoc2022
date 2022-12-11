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

	monkeymap := make(map[int]monkey)

	for i, s := range input {
		lines := strings.Split(s, "\n")
		temp := createMonkey(lines)
		monkeymap[i] = temp
	}

	//Part 1 = false, Part 2 true
	solve(monkeymap, true)

}

func solve(monkeymap map[int]monkey, p2 bool) {
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
				var nextMonkey monkey
				if newLevel%m.test == 0 {
					nextMonkey = monkeymap[m.iftrue]
				} else {
					nextMonkey = monkeymap[m.iffalse]
				}
				nextMonkey = addItems(nextMonkey, newLevel)
				monkeymap[nextMonkey.id] = nextMonkey
				m.items[i] = -1
			}
			m = rearrangeItems(m)
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

func addItems(m monkey, i int) monkey {
	temp := make([]int, 0)
	temp = m.items
	temp = append(temp, i)

	m2 := monkey{}
	m2.id = m.id
	m2.items = temp
	m2.operation = m.operation
	m2.test = m.test
	m2.iftrue = m.iftrue
	m2.iffalse = m.iffalse
	m2.inspected = m.inspected

	return m2

}

func rearrangeItems(m monkey) monkey {
	temp := make([]int, 0)

	m2 := monkey{}
	m2.id = m.id
	m2.operation = m.operation
	m2.test = m.test
	m2.iftrue = m.iftrue
	m2.iffalse = m.iffalse
	m2.inspected = m.inspected

	for _, s := range m.items {
		if s != -1 {
			temp = append(temp, s)
		}
	}
	m2.items = temp
	return m2
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

func createMonkey(lines []string) monkey {
	temp := monkey{}

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
	temp.items = items
	temp.operation = operation
	temp.test = test
	temp.iftrue = iftrue
	temp.iffalse = iffalse
	temp.inspected = 0

	return temp
}

func printMonkey(m monkey) {
	fmt.Println("Monkey", m.id, ":")
	fmt.Print("  Starting Items: ")
	for _, s := range m.items {
		fmt.Print(s, " ")
	}
	fmt.Println("")
	fmt.Println("  Operation: new = old", m.operation)
	fmt.Println("  Test: divisible by", m.test)
	fmt.Println("    If true: throw to monkey", m.iftrue)
	fmt.Println("    If false: throw to monkey", m.iffalse)
	fmt.Println("      Inspected:", m.inspected)
	fmt.Println("")
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
