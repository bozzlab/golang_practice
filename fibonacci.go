package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(y int) int {
	y := 0
	x := 1
	return func(a int) int {
		y += x
		return y
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
