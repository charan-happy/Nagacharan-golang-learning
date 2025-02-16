// A function is a block of code that performs a specific task. It has a name and it is reusable.

package main

import "fmt"

// Function definition
func multiply(a int, b int) int {
	return a * b
}

func main() {
	// Function call
	result := multiply(10, 20)
	fmt.Println(result)
}

/* output
 go run 7-functions.go
200
*/
