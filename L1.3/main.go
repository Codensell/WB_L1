package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	var N int
	for {
		fmt.Println("Сколько воркеров?/how many workers?")
		_, err := fmt.Scan(&N)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Нужно целое число/Need integer")
			continue
		}
		if N <= 0 {
			fmt.Fprintln(os.Stderr, "Количество должно быть больше 0/Number must be > 0")
			continue
		}
		break
	}

	ch := make(chan int, N)

	for i := 0; i < N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for value := range ch {
				fmt.Printf("воркер/worker %d - %d\n", i, value)
			}
		}(i)
	}

	var input int
	for {
		_, err := fmt.Scan(&input)
		if err != nil {
			break
		}
		ch <- input
	}
	close(ch)
	wg.Wait()
}
