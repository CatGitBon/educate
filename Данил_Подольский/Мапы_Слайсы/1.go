package main

import (
	"fmt"
	"sort"
)

func main() {
	v := []int{3, 4, 1, 2, 5} // len = cap
	ap(v)                     //
	sr(v)                     // отсортирует массив
	fmt.Println(v)            // тот же массив
}

func ap(arr []int) {
	arr = append(arr, 10) // сдесь будет новый массив, тк придется выделить память
}

func sr(arr []int) {
	sort.Ints(arr) // а здесь будет сортированный массив и он изменит исходный
}
