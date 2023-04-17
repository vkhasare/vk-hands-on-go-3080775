// challenges/types-primitive/begin/main.go
package main

import "fmt"

// use type inference to create a package-level variable "val" and assign it a value of "global"
var val string = "global"

func printGlobalVal() {
	fmt.Println(val)
}

func updateGlobalVal(newVal string) {
	val = newVal
}

func main() {
	// create a local variable "val" and assign it an int value
	val := 1_00

	// print the value of the local variable "val"
	fmt.Println(val)

	// print the value of the package-level variable "val"
	printGlobalVal()

	// update the package-level variable "val" and print it again
	updateGlobalVal("newGlobalVal")
	printGlobalVal()

	// print the pointer address of the local variable "val"
	ptr := &val
	fmt.Printf("val = %v\n", &val)
	fmt.Printf("ptr = %v\n", &ptr)
	fmt.Printf("*ptr = %v\n", *ptr)
	*ptr = 10_000
	fmt.Printf("*ptr = %v\n", *ptr)

	// assign a value directly to the pointer address of the local variable "val" and print the value of the local variable "val"
}
