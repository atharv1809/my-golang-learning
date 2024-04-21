package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating of our pizza.")
	input, _ := reader.ReadString('\n')
	fmt.Println("Entered data is:", input)
}
