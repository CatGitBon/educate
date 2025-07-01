package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	a := 10
	var c int32 = 0

	wg := &sync.WaitGroup{}

	wg.Add(a)
	for i := 0; i < a; i++ {

		go func() {
			defer wg.Done()
			fmt.Println(i)
			atomic.AddInt32(&c, 1)
		}()

	}

	wg.Wait()

	fmt.Printf("c = %d", c)
}
