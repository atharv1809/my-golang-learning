package main

import "fmt"

func main() {
	fmt.Println("Welcome to Structs in Golang")

	Atharv := User{"Atharv", "atharv@gmail.com", true, 16}
	fmt.Printf("Atharv is:%+v\n", Atharv)

	Atharv.getStatus()
	Atharv.setEmail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) getStatus() {
	fmt.Println("Status of the user is:", u.Status)
}

func (u User) setEmail() {
	u.Email = "bhadange@gmail.com"
	fmt.Println("New Email is:", u.Email)
}
