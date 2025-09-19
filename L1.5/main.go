/*
Разработать программу, которая будет последовательно отправлять значения в
канал, а с другой стороны канала – читать эти значения. По истечении N секунд
программа должна завершаться.
*/

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	var N int
	if _, err := fmt.Scan(&N); err != nil || N <= 0 {
		fmt.Println("need positive number")
		return
	}

	periodTick := 1 * time.Second

	ctx, stopFn := context.WithTimeout(context.Background(), time.Duration(N)*time.Second)
	defer stopFn()

	ch := make(chan int)

	go func() {
		defer close(ch)
		tick := time.NewTicker(periodTick)
		defer tick.Stop()

		i := 0
		for {
			select {
			case <-ctx.Done():
				return
			case <-tick.C:
				ch <- i
				i++
			}
		}

	}()
	for value := range ch {
		fmt.Println(value)
	}
	fmt.Println("Program finished")
}
