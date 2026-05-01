package main

import (
	"fmt"
	// "example.go/mathlib"
)

// import "mathlib"
func sum(a, b int) { // parameter or formal parameter
	z := a + b
	fmt.Println(z)
}

func main() {
	sum(5, 10) //argument or actual parameter
}

func init() {
	fmt.Println("init function called")

	// fmt.Println(a)
}
