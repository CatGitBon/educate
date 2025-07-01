package main

import (
	"fmt"
)

func main() {
	ch := make(chan bool)

	go func() {
		fmt.Println(<-ch)
		fmt.Println(<-ch)

	}()

	ch <- true
	ch <- true

	// time.Sleep(time.Second * 5)
}
