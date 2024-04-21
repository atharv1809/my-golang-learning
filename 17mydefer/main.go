package main

import "fmt"

func main() {
	fmt.Println("Welcome to Defer in golang")
	// defer can be imagined as picking the statement and putting it before closing curly braces of thet function
	// it multiple Defers are there, then the Defer statements are put before curly braces according to LIFO principle
	// without  defer statements are executed with the same sequence the appear

	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")
	fmt.Println("Hello")

	// because of defer the execution will be as
	// it will be like LIFO, latest defered statement will be the first to execute
	// fmt.Println("Three")
	// fmt.Println("Two")
	// fmt.Println("One")
}

func myDefer() {
	var i int
	for i = 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	// this function will finally look like
	// fmt.Println(4)
	// fmt.Println(3)
	// fmt.Println(2)
	// fmt.Println(1)
	// fmt.Println(0)
}
