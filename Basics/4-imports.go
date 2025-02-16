// This code groups the imports into a parenthesized, "factored" import statement.

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	// This demonstrates how Go can combine string formatting with mathematical computations to produce dynamic output.
}

/* output
$ go run imports.go
Now you have 2.6457513110645907 problems.
*/
