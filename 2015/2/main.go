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

	f := l * w
	s := w * h
	t := h * l

	smallest := myMin(f, myMin(s, t))

	return (f+s+t)*2 + smallest
}

func myMin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
