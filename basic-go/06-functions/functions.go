package main

import "fmt"

func main() {
	// Here we call a function that takes two ints and returns their sum as an int.
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	// When you have multiple consecutive parameters of the same type,
	// you may omit the type name for the like-typed parameters up to the final parameter that declares the type.
	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}

// Functions can take zero or more arguments.
func plus(a int, b int) int {
	// Go requires explicit returns, i.e. it won't automatically return the value of the last expression.
	return a + b
}

// When you have multiple consecutive parameters of the same type,
// you may omit the type name for the like-typed parameters up to the final parameter that declares the type.
func plusPlus(a, b, c int) int {
	return a + b + c
}
