package main

import (
	"fmt"
	"time"
)

// Задача 5
// 1. Как будет работать код?
// 2. Как сделать так, чтобы выводился только первый ch?
// ===========================================================

func main() {
	ch := make(chan bool)
	ch2 := make(chan bool)
	ch3 := make(chan bool)
	go func() {
		ch <- true
	}()
	go func() {
		ch2 <- true
	}()
	go func() {
		ch3 <- true
	}()

	go func() {

		for i := 0; i < 3; i++ {
			select {
			case <-ch:
				fmt.Printf("val from ch")
			}

		}
	}()

	time.Sleep(time.Second * 3)
}
