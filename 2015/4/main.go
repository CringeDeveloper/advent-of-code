package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func main() {
	input := "yzbqklnj"
	code := 0

	for {
		codeStr := strconv.Itoa(code)
		result := fmt.Sprintf("%x", md5.Sum([]byte(input+codeStr)))
		if result[:6] == "000000" {
			fmt.Println(codeStr)
			fmt.Println(result)
			return
		}

		code++
	}
}
