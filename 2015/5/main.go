package main

import (
	"fmt"
	"strings"
)

func main() {
	str := ``

	lines := strings.Split(str, "\n")
	good := 0

	for _, v := range lines {
		if stringGood(v) {
			good++
		}
	}

	fmt.Println(good)
}

func stringGood(str string) bool {
	first := ""
	second := ""
	rep := false
	par := false
	pairs := make(map[string]int)

	for i, v := range str {
		if first != "" {
			index, ok := pairs[first+string(v)]
			if ok {
				if index != i-1 {
					par = true
				}
			}
			if !ok {
				pairs[first+string(v)] = i
			}

			if second != "" {
				if string(v) == second {
					rep = true
				}
			}
		}

		if par && rep {
			return true
		}

		second = first
		first = string(v)
	}

	return false
}
