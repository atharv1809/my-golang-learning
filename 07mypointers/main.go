package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on Pointers")

	// var ptr *int
	// fmt.Println("Value of pointer is:", ptr) // ans is <nil>

	myNumber := 23

	var ptr = &myNumber

	fmt.Println("Value of actual pointer is:", ptr)
	fmt.Println("Value of inside pointer is:", *ptr)
}
