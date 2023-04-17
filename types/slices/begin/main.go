// types/slices/begin/main.go
package main

import "fmt"

func main() {
	// declare a slice string the make function
	mySlice := make([]string, 0, 3)

	// append 3 names to the slice
	mySlice = append(mySlice, "John")
	mySlice = append(mySlice, "Jane")
	mySlice = append(mySlice, "Mary")

	// print the slice
	fmt.Println("Current slice", mySlice)

	// slice the slice using syntax slice[low:high]
	fmt.Println("1:3 slice", mySlice[1:3])
	// [Jane Mary]
	// [Jane Mary]
	fmt.Println("1: slice", mySlice[1:])
	// [John]
	fmt.Println(":2 slice", mySlice[:2])
	// [John Jane Mary]
	fmt.Println(": slice", mySlice[:])

	fmt.Println("len", len(mySlice), "capacity", cap(mySlice))

	myNumSlice := make([]int, 3)

	fmt.Println("Initial len", len(mySlice), "capacity", cap(mySlice))

	myNumSlice = append(myNumSlice, 1)
	myNumSlice = append(myNumSlice, 12)
	myNumSlice = append(myNumSlice, 123)
	

	fmt.Println("Current slice", myNumSlice)

	// slice the slice using syntax slice[low:high]
	fmt.Println("1:3 num slice", myNumSlice[1:3])
	
	// [Jane Mary]
	// [Jane Mary]
	fmt.Println("1: num slice", myNumSlice[1:])
	// [John]
	fmt.Println(":2 num slice", myNumSlice[:2])
	// [John Jane Mary]
	fmt.Println(": num slice", myNumSlice[:])
	fmt.Println("len", len(mySlice), "capacity", cap(mySlice))
}
