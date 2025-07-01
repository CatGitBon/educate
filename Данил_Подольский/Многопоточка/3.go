package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		fmt.Println(<-ch)
	}()

	ch <- 1

	time.Sleep(time.Second * 1)
}
