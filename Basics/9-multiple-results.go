package main

import "fmt"

// swap function to demonstrate multiple results
func swap(a, b string) (string, string) {
	return b, a
}

func main() {
	a, b := "hello", "charan"
	fmt.Println("before swap function:", a, b)
	a, b = swap(a, b)
	fmt.Println("after swap function:", a, b)
}

/* output
 go run 9-multiple-results.go
before swap function: hello charan
after swap function: charan hello
*/
