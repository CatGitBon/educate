package main

import (
	"fmt"
	"sync"
)

// Задача 6
// 1. Что выведет код и как исправить?
// ===========================================================

var globalMap = map[string][]int{"test": make([]int, 0), "test2": make([]int, 0), "test3": make([]int, 0)}
var a = 0

func main() {
	wg := sync.WaitGroup{}
	mutex := &sync.Mutex{}

	wg.Add(3)
	go func() {
		defer wg.Done()

		mutex.Lock()
		a = 10
		globalMap["test"] = append(globalMap["test"], a)
		mutex.Unlock()

	}()

	go func() {
		defer wg.Done()
		mutex.Lock()
		a = 11
		globalMap["test2"] = append(globalMap["test2"], a)
		mutex.Unlock()
	}()

	go func() {
		defer wg.Done()
		mutex.Lock()
		a = 12
		globalMap["test3"] = append(globalMap["test3"], a)
		mutex.Unlock()
	}()

	wg.Wait()
	fmt.Printf("%v\n", globalMap)
	fmt.Printf("%d", a)
}
