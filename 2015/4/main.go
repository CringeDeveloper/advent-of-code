package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"sync"
)

type abc struct {
	code int
	mu   sync.Mutex
}

func (a *abc) get() int {
	a.mu.Lock()
	defer a.mu.Unlock()
	code := a.code
	a.code += 1000
	return code
}

func main() {
	input := "yzbqklnj"
	code := 9962624

	semaphore := make(chan struct{}, 5)

	test := abc{-1000, sync.Mutex{}}

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		for {
			semaphore <- struct{}{}

			go func() {
				curr := test.get()
				for i := curr; i < curr+1000; i++ {
					codeStr := strconv.Itoa(i)
					result := fmt.Sprintf("%x", md5.Sum([]byte(input+codeStr)))
					if result[:6] == "000000" {
						fmt.Println(codeStr)
						fmt.Println(result)
						wg.Done()
						return
					}

					code++
				}
				<-semaphore
			}()
		}
	}()

	wg.Wait()
}
