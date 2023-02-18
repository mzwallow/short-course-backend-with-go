package main

import "fmt"

// Go has only one looping construct, the for loop.
func main() {
	// The most basic type, with a single condition.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// A classic initial/condition/after for loop.
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// If you want to loop over a collection of values, use the range
	// form of the for loop. Here we range over a slice of strings
	// and print each index and value.
	strs := []string{"a", "b", "c"}
	for index, value := range strs {
		fmt.Println("index:", index, "value:", value)
	}

	// range on map iterates over key/value pairs.
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// range can also iterate over just the keys of a map.
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// range on strings iterates over Unicode code points.
	// The first value is the starting byte index of the rune and the second
	// the rune itself.
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
