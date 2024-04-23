package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com/todos/1"

func main() {
	fmt.Println("Welcome to handling web requests in Golang")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response type is:%T\n", response)

	defer response.Body.Close()
	databytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println("The content is:", content)
}
