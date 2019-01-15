package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 4}
	b := []int{2, 4, 5}
	c := a - b
	fmt.Println(c)
}
