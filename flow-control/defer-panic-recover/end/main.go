package main

import "fmt"

func cleanup(msg string) {
	fmt.Println(msg)
}
func stub() {
	// defer recovery
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from panicing stub:", err)
		}
	}()

	fmt.Println("running in stub...")

	// panic
	panic("something bad happened in stub")
}
func main() {
	// defer function calls
	defer cleanup("first")
	defer cleanup("second")

	// defer recovery
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from panic:", err)
		}
	}()

	fmt.Println("running in main...")

	// panic
	//panic("something bad happened")

	stub()
	fmt.Println("main continued...")
}
