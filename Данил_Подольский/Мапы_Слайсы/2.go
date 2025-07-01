package main

import (
	"fmt"
)

func main() {
	var foo []int
	var bar []int

	foo = append(foo, 1)
	foo = append(foo, 2)
	foo = append(foo, 3) //[1,2,3]] len 3 cap 4
	bar = append(foo, 4) //[1,2,3,5] здесь будет старый массив
	foo = append(foo, 5) //[1,2,3,5] здесь будет изменение всех массивов

	fmt.Println(foo, bar)
}
