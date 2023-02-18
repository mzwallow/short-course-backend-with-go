package main

import "fmt"

// Error handling is a big topic in Go. We'll cover it in depth in the next section.
// For now, here's a quick example of how to use the `error` interface.
func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		fmt.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("norgate math: square root of negative number %v", f)
	}
	// implementation
	return 42, nil
}

// Output:
// norgate math: square root of negative number -10.23
