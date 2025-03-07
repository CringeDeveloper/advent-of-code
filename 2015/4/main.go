package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func main() {
	input := []byte("yzbqklnj")
	code := 0
	f := 0

	for {
		codeS := strconv.Itoa(code)
		temp := append(input, []byte(codeS)...)
		result := fmt.Sprintf("%x", md5.Sum(temp))
		for _, v := range []rune(result[:6]) {
			if v == '0' {
				f++
			}

			if f == 6 {
				fmt.Println(codeS)
				fmt.Println(result)
				return
			}
		}

		f = 0
		code++
	}
}
