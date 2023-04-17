// types/arrays/begin/main.go
package main

import "fmt"

func main() {
	// declare an array of integers
	var myArr [3]int

	// print the array
	fmt.Println("Array contents", myArr)
	myArr[0] = 9

	// print the array
	fmt.Println("Modified array contents", myArr)

}
