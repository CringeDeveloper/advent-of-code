package main

import "fmt"

type compose struct {
	x int
	y int
}

func main() {
	x := 0
	y := 0
	delivered := 1
	m := make(map[compose]int)
	m[compose{x, y}] = 1

	input := []rune("^v^v^v^v^v")

	for _, v := range input {
		switch v {
		case '^':
			y++
			if _, ok := m[compose{x, y}]; !ok {
				m[compose{x, y}] = 1
				delivered++
			}
		case 'v':
			y--
			if _, ok := m[compose{x, y}]; !ok {
				m[compose{x, y}] = 1
				delivered++
			}
		case '<':
			x--
			if _, ok := m[compose{x, y}]; !ok {
				m[compose{x, y}] = 1
				delivered++
			}
		case '>':
			x++
			if _, ok := m[compose{x, y}]; !ok {
				m[compose{x, y}] = 1
				delivered++
			}
		}
	}

	fmt.Println(delivered)
}
