// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Print the literals
//
//  1. Print a few integer literals
//
//  2. Print a few float literals
//
//  3. Print true and false bool constants
//
//  4. Print your name using a string literal
//
//  5. Print a non-english sentence using a string literal
//
// ---------------------------------------------------------

func main() {
	// integer literal
	fmt.Println(
		-1, 0, 100, 999999,
	)

	// float literal
	fmt.Println(
		-1.00, -5.83,6.89, 99999.54,
	)

	// bool constants
	fmt.Println(
		true, false,
	)

	// string literal - utf-8
	fmt.Println(
		"", // empty string
		"hi",
	)

	fmt.Println(		
		// unicode
		"Güzel bir gün",   //turkish
		"美好的一天", 	//chinese
	)
}
