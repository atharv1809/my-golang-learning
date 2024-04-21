package main

import "fmt"

func main() {
	fmt.Println("Welcome to Structs in Golang")

	Atharv := User{"Atharv", "atharv@gmail.com", true, 16}
	fmt.Printf("Atharv is:%+v", Atharv)
}

type User struct { // in structs if the first letter of variable is capital then it is public else private
	Name   string
	Email  string
	Status bool
	Age    int
}
