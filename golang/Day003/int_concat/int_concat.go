package main

import "fmt"

func main() {
	// Integer operations
	num1 := 10
	num2 := 5

	sum := num1 + num2
	difference := num1 - num2
	product := num1 * num2
	quotient := num1 / num2

	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", difference)
	fmt.Println("Product:", product)
	fmt.Println("Quotient:", quotient)

	// Float operations
	float1 := 3.14
	float2 := 2.0

	sumFloat := float1 + float2
	differenceFloat := float1 - float2
	productFloat := float1 * float2
	quotientFloat := float1 / float2

	fmt.Println("Sum (float):", sumFloat)
	fmt.Println("Difference (float):", differenceFloat)
	fmt.Println("Product (float):", productFloat)
	fmt.Println("Quotient (float):", quotientFloat)
}
