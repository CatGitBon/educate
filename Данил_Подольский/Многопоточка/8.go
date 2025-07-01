package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Задача 8
// 1. Что выведется и как исправить?
// ===========================================================

func main() {
	var counter int32 = 0

	wg := &sync.WaitGroup{}

	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func(i int) {
			defer wg.Done()
			atomic.AddInt32(&counter, 1)
		}(i)
	}

	wg.Wait()
	fmt.Println(counter)
}
