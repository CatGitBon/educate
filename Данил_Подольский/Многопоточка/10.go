package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Задача 10
// 1. Merge n channels
// 2. Если один из входных каналов закрывается, то нужно закрыть все остальные каналы
// ===========================================================

//  !!! Научиться исползовать контекст

func case3(wg *sync.WaitGroup, ctx context.Context, cancel context.CancelFunc, channels ...chan int) chan int {

	out := make(chan int)

	wg.Add(len(channels))
	for _, channal := range channels {
		go func(channal chan int) {

			defer wg.Done()

			for {
				select {

				case <-ctx.Done():
				case val, ok := <-channal:
					if !ok {
						cancel()
						return
					}

					out <- val

				}

			}
		}(channal)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	go func() {

		for i := 0; i < 10000; i++ {
			select {
			case <-ctx.Done():
			case ch1 <- i:
			}

		}

		close(ch1)
	}()

	go func() {

		for i := 0; i < 100000; i++ {
			select {
			case <-ctx.Done():
			case ch2 <- i:
			}
		}

		close(ch2)
	}()

	go func() {

		for i := 0; i < 200; i++ {
			select {
			case <-ctx.Done():
			case ch3 <- i:
			}
		}

		close(ch3)
	}()

	out := case3(wg, ctx, cancel, ch1, ch2, ch3)
	for i := range out {
		fmt.Println(i)
	}

}
