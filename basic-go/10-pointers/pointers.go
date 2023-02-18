package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// The &i syntax gives the memory address of i, i.e. a pointer to i.
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// Pointers can be printed too.
	fmt.Println("pointer:", &i)

	// Pointers hold the memory address of a value.
	// The type *T is a pointer to a T value. Its zero value is nil.
	// The & operator generates a pointer to its operand.
	// The * operator denotes the pointer's underlying value.
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	// Pointers are useful for passing references to large objects.
	// For example, a function that takes a slice of bytes can avoid copying the bytes by taking a pointer to the slice.
	// This is more efficient if the slice is large, as it avoids the allocation and copying of the bytes.
}
