/*
Для корректного завершения программы без потери данных используя ctrl+c необходимо использовать signal.NotifyContext.
Возвращает context.Context (здесь background для main функции) и context.CancelFunc (os.Interrupt == Ctrl + C).

При выходе через Ctrl + C вызывается ctx.Done и вызывается wg.Done.
В продюссере так же вызывается ctx.Done и далее close jobs.
Далее вызывается wg.Wait и stop
*/

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	ctx, stopFn := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stopFn()
	jobs := make(chan int)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case _, ok := <-jobs:
					if !ok {
						return
					}
				}
			}
		}(i)
	}

	go func() {
		defer close(jobs)
		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				return
			case jobs <- i:
				time.Sleep(3 * time.Second)
			}
		}
	}()
	fmt.Printf("%T\n%T\n", ctx, stopFn) // просто показать что присвоено переменным
	fmt.Println("Ctrl + C to escape")
	wg.Wait()
}
