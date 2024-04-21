package main

import "fmt"

func main() {
	fmt.Println("Welcome to Structs in Golang")

	Atharv := User{"Atharv", "atharv@gmail.com", true, 16}
	fmt.Printf("Atharv is:%+v", Atharv)
}

type User struct {
	Name   string
	Email  string
	status bool
	Age    int
}
