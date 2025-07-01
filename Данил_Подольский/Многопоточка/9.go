package main

import (
	"fmt"
	"sync"
)

// ===========================================================
// Задача 9
// 1. Что выведется и как исправить?
// 2. Что поправить, чтобы сохранить порядок?
// ===========================================================

func main() {
	wg := &sync.WaitGroup{}
	m := make(chan struct {
		id int
		st string
	}, 3)
	cnt := 5
	sl := make(map[int]string)
	// mutex := &sync.Mutex{}

	wg.Add(cnt)
	for i := 0; i < cnt; i++ {
		go func(i int) {
			defer wg.Done()

			m <- struct {
				id int
				st string
			}{id: i, st: fmt.Sprintf("Goroutine %d", i)}
		}(i)
	}

	wg.Add(cnt)
	for i := 0; i < cnt; i++ {
		go func() {
			defer wg.Done()
			val := <-m

			sl[val.id] = val.st
		}()
	}

	wg.Wait()
	close(m)

	for _, val := range sl {
		fmt.Println(val)
	}
}

// func ReceiveFromCh(ch chan struct {
// 	id int
// 	st string
// }, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	val := <-ch

// 	fmt.Println(val.id)
// }
