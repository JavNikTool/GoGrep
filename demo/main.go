package main

import "fmt"

// TODO: refactor greeting
func main() {
	fmt.Println("hello")

	// FIXME: handle errors properly
	if 2+2 == 4 {
		fmt.Println("OK")
	}

	fmt.Println("operation FAILED") // intentionally uppercase word
}
