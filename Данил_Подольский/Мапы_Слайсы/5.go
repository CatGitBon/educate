package main

import "fmt"

func main() {
	mutate := func(a []int) {
		a[0] = 0         // {0,2,3,4}
		a = append(a, 1) //{0,2,3,4,1}
		fmt.Println(a)   //{0,2,3,4,1}
	}
	a := []int{1, 2, 3, 4}
	mutate(a)
	fmt.Println(a) // //{0,2,3,4}
}
