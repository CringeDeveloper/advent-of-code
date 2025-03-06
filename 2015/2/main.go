package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("1input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	totalSize := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		totalSize += getWrapperSize(scanner.Text())
	}

	fmt.Println(totalSize)
}

func getWrapperSize(size string) int {
	dimensions := strings.Split(size, "x")
	l, _ := strconv.Atoi(dimensions[0])
	w, _ := strconv.Atoi(dimensions[1])
	h, _ := strconv.Atoi(dimensions[2])

	smallest := min(l, min(w, h))
	secondSmallest := max(min(l, w), max(min(w, h), min(h, l)))

	ribbon := (smallest + secondSmallest) * 2

	return (l * w * h) + ribbon
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
