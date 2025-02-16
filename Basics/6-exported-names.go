/* exported names concept in go

- exported names in go are the ones that start with a capital letter
- when you import a package, you can only refer to its exported names
- any "unexported" names are not accessible from outside the package
- for example, fmt.Println is an exported name from the fmt package
- but fmt.print is not an exported name
*/

package main

import "fmt"

func main() {
	fmt.Print("Hello, World!")
}

// here if we try with (small p instead of capital P) fmt.print it will give an error  undefined: fmt.print

/*  output

$ go run 6-exported-names.go
Hello, World!

*/
