package main

import (
	"fmt"
	"runtime"
)

// ---------------------------------------------------------
// EXERCISE: Print the Go Version
//
//  1. Look at the runtime package documentation
//  2. Find the func that returns the Go version
//  3. Print the Go version by calling that func
//
// HINT
//  It's here: https://golang.org/pkg/runtime
//
// EXPECTED OUTPUT
//  "go1.10"
// ---------------------------------------------------------

func main() {
	fmt.Println(runtime.Version()) // output : go1.18.1

	// Version returns the Go tree's version string.
}
