package main

import (
	"fmt"
	// "example.go/mathlib"
)

// import "mathlib"

func processOperation(a int, b int, op func(p, q int)) {
	op(a, b)

}

func call() func(a, b int) {
	return sum

}
func sum(a, b int) { // parameter or formal parameter
	z := a + b
	fmt.Println(z)
}

func main() {
	//argument or actual parameter

	add := call()
	add(10, 20)
}
