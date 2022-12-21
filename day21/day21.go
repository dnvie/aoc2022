package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	i, _ := os.ReadFile("input.txt")
	input := strings.Split(string(i), "\n")
	monkeymap := make(map[string]*Monkey)

	vmap := make(map[string]int)
	rootSet := false

	for !rootSet {
		for i := range input {
			f := strings.Fields(input[i])
			if len(f) == 2 {
				v, _ := strconv.Atoi(f[1])
				vmap[f[0][:len(f[0])-1]] = v
			} else {
				v, b := eval(f[1], f[2], f[3], vmap)
				if b {
					vmap[f[0][:len(f[0])-1]] = v
					if f[0] == "root:" {
						rootSet = true
						break
					}
				}
			}
		}
	}

	for i := range input {
		f := strings.Fields(input[i])
		m := &Monkey{id: f[0][:len(f[0])-1]}
		monkeymap[f[0][:len(f[0])-1]] = m
	}

	for i := range input {
		f := strings.Fields(input[i])
		if len(f) == 2 {
			v, _ := strconv.Atoi(f[1])
			m := monkeymap[f[0][:len(f[0])-1]]
			m.value = int64(v)
			m.leaf = true
		} else {
			m := monkeymap[f[0][:len(f[0])-1]]
			m.left = monkeymap[f[1]]
			monkeymap[f[1]].parent = m
			m.right = monkeymap[f[3]]
			monkeymap[f[3]].parent = m
			m.operation = f[2]
		}
	}
	root := monkeymap["root"]
	root.parent, root.operation = &Monkey{}, "=="
	res := toRoot(monkeymap["humn"], vmap, root.parent)
	fmt.Println("Part 1:", vmap["root"])
	fmt.Println("Part 2: (Paste equation into SageMath and solve for x)")
	fmt.Println(res[1 : len(res)-1])
	fmt.Println("\nTime elapsed:", time.Since(start))

}

func eval(s1 string, s2 string, s3 string, vmap map[string]int) (int, bool) {
	i1, i2, res := 0, 0, 0
	if _, ok := vmap[s1]; ok {
		i1 = vmap[s1]
	} else {
		return 0, false
	}

	if _, ok := vmap[s3]; ok {
		i2 = vmap[s3]
	} else {
		return 0, false
	}

	if s2 == "+" {
		res = i1 + i2
	} else if s2 == "-" {
		res = i1 - i2
	} else if s2 == "*" {
		res = i1 * i2
	} else {
		res = i1 / i2
	}

	return res, true
}

func toRoot(m *Monkey, vmap map[string]int, root *Monkey) string {
	var val int64
	s := "x"

	for m.parent != root {
		if m == m.parent.left {
			val = int64(vmap[m.parent.right.id])
			s = fmt.Sprintf("(%s%s%d)", s, m.parent.operation, val)
		} else {
			val = int64(vmap[m.parent.left.id])
			s = fmt.Sprintf("(%d%s%s)", val, m.parent.operation, s)
		}
		m = m.parent
	}
	return s
}

type Monkey struct {
	id                  string
	left, right, parent *Monkey
	operation           string
	value               int64
	leaf                bool
}
