package main

import "fmt"

func main() {
	welcomeMessage()

	res := add(5, 9) // normal function
	fmt.Println("Addition result is:", res)

	proRes, proMessage := proAdd(5, 9, 6, 1, 3)
	fmt.Println("Pro message is:", proMessage)
	fmt.Println("Pro result is:", proRes)
}

func welcomeMessage() {
	fmt.Println("Welcome to Functions in Golang")
}

func add(val1 int, val2 int) int {
	return val1 + val2
}

func proAdd(values ...int) (int, string) {
	var add int
	for _, val := range values {
		add += val
	}
	return add, "Written a Pro function"
}
