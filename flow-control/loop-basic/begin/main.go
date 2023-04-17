// flow-control/loop-basic/begin/main.go
package main

import "fmt"

func main() {
	// declare a string to iterate over
	myString := "LoveyDoveyCoochieKoo"

	// iterate over the string with basic for loop
	for _, char := range myString {
		fmt.Println("%s", string(char))
	}
}
