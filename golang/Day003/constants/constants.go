package main

import "fmt"

func main() {
	// Declare a constant
	const pi float64 = 3.14159265359

	fmt.Println("Value of pi:", pi)

	// Constants can be used in arithmatic expressions
	radius := 5
	circumference := 2 * pi * float64(radius)
	fmt.Println("Circumference:", circumference)

	// Constants cannot be reassigned
	// pi = 3.14 // This will result in a compilation error
}
