/*
- These format specifiers are part of the fmt package in Go
- These format specifiers allow you to control how different types of data are represented as strings, making it easier to produce readable and well-formatted output
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Integer: %d\n", 42)
	fmt.Printf("Float: %f\n", 3.14159)
	fmt.Printf("Scientific: %e\n", 3.14159)
	fmt.Printf("String: %s\n", "Hello, World!")
	fmt.Printf("Boolean: %t\n", true)
	fmt.Printf("Value: %v\n", 42)
	fmt.Printf("Type: %T\n", 42)
}

/* output
$ go run format-specifiers.go
Integer: 42
Float: 3.141590
Scientific: 3.141590e+00
String: Hello, World!
Boolean: true
Value: 42
Type: int
*/

// The %d format specifier is used to represent integers, %f is used to represent floating-point numbers, %e is used to represent scientific notation, %s is used to represent strings, %t is used to represent booleans, %v is used to represent values, and %T is used to represent types. The %p format specifier is used to represent pointers.
