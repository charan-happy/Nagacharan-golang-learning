// NOTE :  main should not be declared twice in different files in the same folder. I moved file outside the folder and executed the program. It worked fine.

// a go program to know about the current time
package main

import (
	"fmt"
	"time"
)

// here i am using timeMain as the function name due to same folder having another main function in `hello.go“
func main() {
	fmt.Println("let's begin go learning journey")

	fmt.Println("the time now is:", time.Now())
}

// To fix the issue of the main function being redeclared, you can rename the main function in one of the files to avoid the conflict.

/* output ;

$ go run time.go
let's begin go learning journey
the time now is: 2025-02-16 08:10:25.1042678 +0530 IST m=+0.004108201
*/
