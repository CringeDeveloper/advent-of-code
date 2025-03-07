package main

import "fmt"

type compose struct {
	x int
	y int
}

func main() {
	santaX := 0
	santaY := 0
	robotX := 0
	robotY := 0
	delivered := 1
	m := make(map[compose]int)
	m[compose{santaX, santaY}] = 1
	turn := 0

	input := []rune("^v^v^v^v^v")

	for _, v := range input {
		if turn%2 == 0 {
			switch v {
			case '^':
				santaY++
				if _, ok := m[compose{santaX, santaY}]; !ok {
					m[compose{santaX, santaY}] = 1
					delivered++
				}
			case 'v':
				santaY--
				if _, ok := m[compose{santaX, santaY}]; !ok {
					m[compose{santaX, santaY}] = 1
					delivered++
				}
			case '<':
				santaX--
				if _, ok := m[compose{santaX, santaY}]; !ok {
					m[compose{santaX, santaY}] = 1
					delivered++
				}
			case '>':
				santaX++
				if _, ok := m[compose{santaX, santaY}]; !ok {
					m[compose{santaX, santaY}] = 1
					delivered++
				}
			}
		} else {
			switch v {
			case '^':
				robotY++
				if _, ok := m[compose{robotX, robotY}]; !ok {
					m[compose{robotX, robotY}] = 1
					delivered++
				}
			case 'v':
				robotY--
				if _, ok := m[compose{robotX, robotY}]; !ok {
					m[compose{robotX, robotY}] = 1
					delivered++
				}
			case '<':
				robotX--
				if _, ok := m[compose{robotX, robotY}]; !ok {
					m[compose{robotX, robotY}] = 1
					delivered++
				}
			case '>':
				robotX++
				if _, ok := m[compose{robotX, robotY}]; !ok {
					m[compose{robotX, robotY}] = 1
					delivered++
				}
			}
		}

		turn++
	}

	fmt.Println(delivered)
}
