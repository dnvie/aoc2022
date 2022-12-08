// Part 1
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	i, _ := os.ReadFile("input.txt")
	input := strings.Split(string(i), "\n")[1:]

	dirs := []*dir{}
	root := &dir{name: "/"}
	root.parent = root
	c := root
	dirs = append(dirs, root)

	for _, s := range input {
		a := strings.Fields(s)

		if a[0] == "$" {
			if a[1] == "cd" {
				if a[2] == ".." {
					c = c.parent
				} else {
					c = c.subs[get(c.subs, a[2])]
				}
			}
		} else if a[0] == "dir" {
			temp := &dir{name: a[1], parent: c}
			c.subs = append(c.subs, temp)
			dirs = append(dirs, temp)
		} else {
			c.files = append(c.files, a[0])
			size, _ := strconv.Atoi(a[0])
			c.size += size
		}
	}

	res1 := 0
	for _, s := range dirs {
		if sizes(s) <= 100000 {
			res1 += sizes(s)
		}
	}

	fmt.Println("Part 1:", res1)
}

func sizes(d *dir) int {
	i := d.size
	for _, s := range d.subs {
		if len(s.subs) == 0 {
			i += s.size
		} else {
			i += sizes(s)
		}
	}
	return i
}

func print(d *dir) {
	fmt.Println("Name:", d.name)
	fmt.Println("Parent:", d.parent.name)
	fmt.Print("Subs:")
	for i := range d.subs {
		fmt.Print(" ", d.subs[i].name)
	}
	fmt.Print("\nFiles:")
	for i := range d.files {
		fmt.Print(" ", d.files[i])
	}
	fmt.Println("")
	fmt.Println("Size:", d.size)
}

func get(d []*dir, s string) int {
	for i, x := range d {
		if x.name == s {
			return i
		}
	}
	return -1
}

type dir struct {
	parent *dir
	name   string
	subs   []*dir
	files  []string
	size   int
}
