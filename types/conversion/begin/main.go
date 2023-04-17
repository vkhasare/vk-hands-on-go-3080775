// types/conversion/begin/main.go
package main

import "fmt"

func main() {
	// declare float and convert it to an unsigned int
	var a float32 = 3.14
	b := int(a)
	fmt.Printf("%v=%T, %v=%T\n", a, a, b, b)
}
