package main

import "fmt"

func main() {
	c := []string{"A", "B", "D", "E"}
	b := c[1:2]         // "B"
	b = append(b, "TT") // "B" "TT"
	fmt.Println(c)      // "A", "B" "TT", "E"
	fmt.Println(b)      // "B" "TT"
}
