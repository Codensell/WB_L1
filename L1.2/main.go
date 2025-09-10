package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	array := []int{2, 4, 6, 8, 10}
	ch := make(chan int)

	go func() {
		wg.Wait()
		close(ch)
	}()

	for _, numArray := range array {
		wg.Add(1)
		go func(elem int) {
			defer wg.Done()
			ch <- elem * elem
		}(numArray)
	}

	for sqr := range ch {
		fmt.Println(sqr)
	}
}
